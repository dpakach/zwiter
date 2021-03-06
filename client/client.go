package client

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/dpakach/zwiter/posts/postspb"
	"github.com/dpakach/zwiter/users/userspb"
	"golang.org/x/oauth2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/testdata"
)

func getAddr(service string) string {
	switch service {
	case "users":
		host := os.Getenv("USER_HOST")
		port := os.Getenv("USER_PORT")
		if host == "" {
			host = "127.0.0.1"
		}
		if port == "" {
			port = "8002"
		}
		return host + ":" + port
	case "posts":
		host := os.Getenv("POST_HOST")
		port := os.Getenv("POST_PORT")
		if host == "" {
			host = "127.0.0.1"
		}
		if port == "" {
			port = "8001"
		}
		return host + ":" + port
	}
	return ""
}

// fetchToken simulates a token lookup and omits the details of proper token
// acquisition. For examples of how to acquire an OAuth2 token, see:
// https://godoc.org/golang.org/x/oauth2
func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "some-secret-token",
	}
}

func getOpts() []grpc.DialOption {
	// Set up the credentials for the connection.
	perRPC := oauth.NewOauthAccess(fetchToken())
	creds, err := credentials.NewClientTLSFromFile(testdata.Path("ca.pem"), "x.test.youtube.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	opts := []grpc.DialOption{
		// In addition to the following grpc.DialOption, callers may also use
		// the grpc.CallOption grpc.PerRPCCredentials with the RPC invocation
		// itself.
		// See: https://godoc.org/google.golang.org/grpc#PerRPCCredentials
		grpc.WithPerRPCCredentials(perRPC),
		// oauth.NewOauthAccess requires the configuration of transport
		// credentials.
		grpc.WithTransportCredentials(creds),
	}
	return opts
}

// NewPostsClient creates new client for Post service
func NewPostsClient() (*grpc.ClientConn, postspb.PostsServiceClient) {
	addr := getAddr("posts")
	opts := getOpts()
	cc, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatal(err)
	}
	return cc, postspb.NewPostsServiceClient(cc)
}

// NewUsersClient creates a new client for the users service
func NewUsersClient() (*grpc.ClientConn, userspb.UsersServiceClient) {
	addr := getAddr("users")
	opts := getOpts()
	cc, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatal(err)
	}
	return cc, userspb.NewUsersServiceClient(cc)
}

// CreatePost create a new post
func CreatePost(c postspb.PostsServiceClient, content string) []byte {
	req := &postspb.CreatePostRequest{
		Text:     content,
		AuthorId: 1,
	}

	res, err := c.CreatePost(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	return getJSON(res)
}

// GetPosts Get all posts
func GetPosts(c postspb.PostsServiceClient) []byte {
	req := &postspb.EmptyData{}

	res, err := c.GetPosts(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	return getJSON(res)
}

// GetPost Get one post by Id
func GetPost(c postspb.PostsServiceClient, id int64) []byte {
	req := &postspb.GetPostRequest{
		Id: id,
	}
	res, err := c.GetPost(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	return getJSON(res)
}

// CreateUser create a new user
func CreateUser(c userspb.UsersServiceClient, username string) []byte {
	req := &userspb.CreateUserRequest{
		Username: username,
	}

	res, err := c.CreateUser(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	return getJSON(res)
}

// GetUsers Get all users
func GetUsers(c userspb.UsersServiceClient) []byte {
	req := &userspb.EmptyData{}

	res, err := c.GetUsers(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	return getJSON(res)
}

// GetUser Get one user by Id
func GetUser(c userspb.UsersServiceClient, id int64) []byte {
	req := &userspb.GetUserRequest{
		Id: id,
	}
	res, err := c.GetUser(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	return getJSON(res)
}

func getJSON(res interface{}) []byte {
	json, err := json.MarshalIndent(res, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	return json
}
