package client

import (
	"context"
	"fmt"
	"log"

	"github.com/dpakach/zwiter/posts/postspb"
	"github.com/dpakach/zwiter/users/userspb"

	"google.golang.org/grpc"
)

// NewPostsClient creates new client for Post service
func NewPostsClient() (*grpc.ClientConn, postspb.PostsServiceClient) {
	cc, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	return cc, postspb.NewPostsServiceClient(cc)
}

// NewUsersClient creates a new client for the users service
func NewUsersClient() (*grpc.ClientConn, userspb.UsersServiceClient) {
	cc, err := grpc.Dial("localhost:8002", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return cc, userspb.NewUsersServiceClient(cc)
}

// CreatePost create a new post
func CreatePost(c postspb.PostsServiceClient) {
	req := &postspb.CreatePostRequest{
		Text:     "This is a sample Post",
		AuthorId: 1,
	}

	res, err := c.CreatePost(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	log.Printf("Respnse from Server: %v", res)
}

// GetPosts Get all posts
func GetPosts(c postspb.PostsServiceClient) {
	req := &postspb.EmptyData{}

	res, err := c.GetPosts(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Respnse from Server: %v", res)
}

// GetPost Get one post by Id
func GetPost(c postspb.PostsServiceClient) {
	req := &postspb.GetPostRequest{
		Id: 1,
	}
	res, err := c.GetPost(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Respnse from Server: %v", res)
}

// CreateUser create a new user
func CreateUser(c userspb.UsersServiceClient) {
	req := &userspb.CreateUserRequest{
		Username: "SampleUser",
	}

	res, err := c.CreateUser(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response from Server: %v", res)
}

// GetUsers Get all users
func GetUsers(c userspb.UsersServiceClient) {
	req := &userspb.EmptyData{}

	res, err := c.GetUsers(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Respnse from Server: %v", res)
}

// GetUser Get one user by Id
func GetUser(c userspb.UsersServiceClient) {
	req := &userspb.GetUserRequest{
		Id: 1,
	}
	res, err := c.GetUser(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Respnse from Server: %v", res)
}
