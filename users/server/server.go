package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dpakach/zwiter/store"
	"github.com/dpakach/zwiter/users/userspb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/testdata"
)

// UserStore is database Instance of users
var UserStore = store.New("users")

// Users type for Users List
type Users struct {
	Users []User `json:"users"`
}

// AddDbList adds a User type to Users.Users
func (p *Users) AddDbList(user *User) {
	p.Users = append(p.Users, *user)
}

// ReadFromDb Updates the database Instance from the file
func (p *Users) ReadFromDb() {
	content := UserStore.GetContent()
	var users Users
	err := json.Unmarshal(content, &users)
	p.Users = users.Users
	if err != nil {
		log.Fatalf("Error while reading db: %v", err)
	}
}

// CommitDb writes the current database instance to the file
func (p *Users) CommitDb() {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	UserStore.Write(jsonData)
}

// GetByID returns a user with given ID from the database instance
func (p *Users) GetByID(id int64) *User {
	p.ReadFromDb()
	for _, item := range p.Users {
		if item.ID == int64(id) {
			return &item
		}
	}
	return nil
}

// NewID returns a new ID for creating new database object
func (p *Users) NewID() int64 {
	id := int64(0)
	for _, user := range p.Users {
		if user.ID > id {
			id = user.ID
		}
	}
	return id + 1
}

// User type for a single user instance
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Created  int64  `json:"created"`
}

// GetID returns ID of the user
func (p *User) GetID() int64 { return p.ID }

// SetID sets ID of the user
func (p *User) SetID(id int64) { p.ID = id }

// SaveToStore saves the user to the store
func (p *User) SaveToStore(store *store.Store) int64 {
	users := new(Users)
	users.ReadFromDb()
	id := p.GetID()
	if id == 0 {
		id = users.NewID()
	}
	p.SetID(id)
	users.AddDbList(p)
	users.CommitDb()
	return id
}

type server struct{}

func (s *server) CreateUser(ctx context.Context, req *userspb.CreateUserRequest) (*userspb.CreateUserResponse, error) {
	ts := time.Now().Unix()
	user := User{Username: req.GetUsername(), Created: int64(ts)}
	id := user.SaveToStore(UserStore)

	created := &Users{}
	created.ReadFromDb()
	createdUser := created.GetByID(id)
	return &userspb.CreateUserResponse{
		Id:       createdUser.ID,
		Username: createdUser.Username,
		Created:  createdUser.Created,
	}, nil
}

func (s *server) GetUsers(ctx context.Context, req *userspb.EmptyData) (*userspb.GetUsersResponse, error) {
	users := new(Users)
	users.ReadFromDb()
	resp := []*userspb.GetUsersResponse_User{}
	for _, user := range (*users).Users {
		resp = append(resp, &userspb.GetUsersResponse_User{
			Id:       user.ID,
			Username: user.Username,
			Created:  user.Created,
		})
	}
	return &userspb.GetUsersResponse{
		Users: resp,
	}, nil
}

func (s *server) GetUser(ctx context.Context, req *userspb.GetUserRequest) (*userspb.GetUserResponse, error) {
	users := new(Users)
	users.ReadFromDb()
	user := users.GetByID(req.GetId())
	return &userspb.GetUserResponse{
		Id:       user.ID,
		Username: user.Username,
		Created:  user.Created,
	}, nil
}

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

// valid validates the authorization.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	return token == "some-secret-token"
}

// ensureValidToken ensures a valid token exists within a request's metadata. If
// the token is missing or invalid, the interceptor blocks execution of the
// handler and returns an error. Otherwise, the interceptor invokes the unary
// handler.
func ensureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	// The keys within metadata.MD are normalized to lowercase.
	// See: https://godoc.org/google.golang.org/grpc/metadata#New
	if !valid(md["authorization"]) {
		return nil, errInvalidToken
	}
	// Continue execution of handler after ensuring a valid token.
	return handler(ctx, req)
}

func getOpts() []grpc.ServerOption {
	cert, err := tls.LoadX509KeyPair(testdata.Path("server1.pem"), testdata.Path("server1.key"))
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}
	opts := []grpc.ServerOption{
		// The following grpc.ServerOption adds an interceptor for all unary
		// RPCs. To configure an interceptor for streaming RPCs, see:
		// https://godoc.org/google.golang.org/grpc#StreamInterceptor
		grpc.UnaryInterceptor(ensureValidToken),
		// Enable TLS for all incoming connections.
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	return opts
}

func main() {
	log.Println("Starting Users Server")

	// Set the filepath for the json store
	if len(os.Args) < 2 {
		log.Fatal(fmt.Errorf("Opps, Seems like you forgot to provide the path of the store file"))
		os.Exit(1)
	}
	path := os.Args[1]
	abs, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(fmt.Errorf("Invalid path provided, Make sure the path %q is correct", path))
		os.Exit(1)
	}
	_, err = os.Stat(abs)
	if os.IsNotExist(err) {
		log.Fatal(fmt.Errorf("Error, Make sure the path %q exists", path))
		os.Exit(1)
	}
	log.Printf("Setting Store file path to: %v\n", abs)
	UserStore.SetFilePath(abs)

	log.Println("Starting Posts server on port 8002")
	lis, err := net.Listen("tcp", "0.0.0.0:8002")
	if err != nil {
		log.Fatal(err)
	}
	opts := getOpts()
	s := grpc.NewServer(opts...)

	userspb.RegisterUsersServiceServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
