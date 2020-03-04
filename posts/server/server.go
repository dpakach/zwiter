package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dpakach/zwiter/client"
	"github.com/dpakach/zwiter/posts/postspb"
	"github.com/dpakach/zwiter/store"
	"github.com/dpakach/zwiter/users/userspb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/testdata"
)

// Posts type for Users List
type Posts struct {
	Posts []Post `json:"posts"`
}

// AddDbList adds a new Post to Posts.Posts
func (p *Posts) AddDbList(post *Post) {
	p.Posts = append(p.Posts, *post)
}

// ReadFromDb updates Posts.Posts from the file
func (p *Posts) ReadFromDb() {
	content := PostStore.GetContent()
	var posts Posts
	err := json.Unmarshal(content, &posts)
	p.Posts = posts.Posts
	if err != nil {
		log.Fatalf("Error while reading db: %v", err)
	}
}

// CommitDb writes Posts.Posts to the file
func (p *Posts) CommitDb() {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	PostStore.Write(jsonData)
}

// GetByID returns Post with given ID
func (p *Posts) GetByID(id int64) *Post {
	p.ReadFromDb()
	for _, item := range p.Posts {
		if item.ID == int64(id) {
			return &item
		}
	}
	return nil
}

// NewID returns new ID for the new Post
func (p *Posts) NewID() int64 {
	id := int64(0)
	for _, post := range p.Posts {
		if post.ID > id {
			id = post.ID
		}
	}
	return id + 1
}

// Post type for a Post Instance
type Post struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Created int64  `json:"created"`
	Author  int64  `json:"author"`
}

// GetID returns ID of the Post
func (p *Post) GetID() int64 { return p.ID }

// SetID sets ID of the Post
func (p *Post) SetID(id int64) { p.ID = id }

// SaveToStore saves the Post to the given dB store
func (p *Post) SaveToStore(store *store.Store) int64 {
	posts := new(Posts)
	posts.ReadFromDb()
	id := p.GetID()
	if id == 0 {
		id = posts.NewID()
	}
	p.SetID(id)
	posts.AddDbList(p)
	posts.CommitDb()
	return id
}

// PostStore is database Instance of posts
var PostStore = store.New("posts")

type server struct{}

func (s *server) CreatePost(ctx context.Context, req *postspb.CreatePostRequest) (*postspb.CreatePostResponse, error) {
	ts := time.Now().Unix()
	post := Post{Title: req.GetText(), Created: int64(ts)}
	userid, ok := ctx.Value("userid").(int64)
	if !ok {
		return nil, errors.New("Invalid userid")
	}
	user, err := getUserByID(userid)
	if err != nil {
		return nil, err
	}


	post.Author = userid
	id := post.SaveToStore(PostStore)
	created := &Posts{}
	created.ReadFromDb()
	createdPost := created.GetByID(id)

	return &postspb.CreatePostResponse{
		Id:       createdPost.ID,
		Text:     createdPost.Title,
		Created:  createdPost.Created,
		AuthorId: user.Id,
	}, nil
}

func getUserByID(id int64) (*userspb.GetUserResponse, error) {
	cc, c := client.NewUsersClient()
	defer cc.Close()
	userReq := &userspb.GetUserRequest{Id: id}
	user, err := c.GetUser(context.Background(), userReq)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *server) GetPosts(ctx context.Context, req *postspb.EmptyData) (*postspb.GetPostsResponse, error) {
	posts := new(Posts)
	posts.ReadFromDb()
	resp := []*postspb.GetPostsResponse_Post{}
	for _, post := range (*posts).Posts {
		user, err := getUserByID(post.Author)
		if err != nil {
			return nil, err
		}

		resp = append(resp, &postspb.GetPostsResponse_Post{
			Id:      post.ID,
			Text:    post.Title,
			Created: post.Created,
			Author: &postspb.GetPostsResponse_User{
				Username: user.Username,
				Id:       user.Id,
			},
		})
	}
	return &postspb.GetPostsResponse{
		Posts: resp,
	}, nil
}

func (s *server) GetPost(ctx context.Context, req *postspb.GetPostRequest) (*postspb.GetPostResponse, error) {
	posts := new(Posts)
	posts.ReadFromDb()
	post := posts.GetByID(req.GetId())
	user, err := getUserByID(post.Author)
	if err != nil {
		return nil, err
	}

	return &postspb.GetPostResponse{
		Id:      post.ID,
		Text:    post.Title,
		Created: post.Created,
		Author: &postspb.GetPostResponse_User{
			Username: user.Username,
			Id:       user.Id,
		},
	}, nil
}

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

// valid validates the authorization.
func valid(authorization []string) (bool, int64) {
	if len(authorization) < 1 {
		return false, 0
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	cc, ac := client.NewAuthClient()
	defer cc.Close()
	res := client.ValidateToken(ac, token)

	return res.Valid, res.GetUserid()
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
	isValid, id := valid(md["authorization"])
	if !isValid {
		return nil, errInvalidToken
	}
	c := context.WithValue(ctx, "userid", id)
	// Continue execution of handler after ensuring a valid token.
	return handler(c, req)
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
	log.Println("Starting Posts Server")
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
	PostStore.SetFilePath(abs)

	log.Println("Starting Posts server on port 8001")
	lis, err := net.Listen("tcp", "0.0.0.0:8001")
	if err != nil {
		log.Fatal(err)
	}
	opts := getOpts()
	s := grpc.NewServer(opts...)

	postspb.RegisterPostsServiceServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
