package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/dpakach/zwiter/store"
	"github.com/dpakach/zwiter/users/userspb"

	"google.golang.org/grpc"
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
	s := grpc.NewServer()

	userspb.RegisterUsersServiceServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
