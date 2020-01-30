package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/dpakach/zwiter/auth/authpb"
	"github.com/dpakach/zwiter/store"

	"google.golang.org/grpc"
)

// Posts type for Users List
type Tokens struct {
	Tokens []Token `json:"tokens"`
}

// AddDbList adds a new Post to Posts.Posts
func (p *Tokens) AddDbList(post *Token) {
	p.Tokens = append(p.Tokens, *post)
}

// ReadFromDb updates Posts.Posts from the file
func (p *Tokens) ReadFromDb() {
	content := TokenStore.GetContent()
	var posts Tokens
	err := json.Unmarshal(content, &posts)
	p.Tokens = posts.Tokens
	if err != nil {
		log.Fatalf("Error while reading db: %v", err)
	}
}

// CommitDb writes Posts.Posts to the file
func (p *Tokens) CommitDb() {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	TokenStore.Write(jsonData)
}

// GetByID returns Post with given ID
func (p *Tokens) GetByID(id int64) *Token {
	p.ReadFromDb()
	for _, item := range p.Tokens {
		if item.ID == int64(id) {
			return &item
		}
	}
	return nil
}

// NewID returns new ID for the new Post
func (p *Tokens) NewID() int64 {
	id := int64(0)
	for _, post := range p.Tokens {
		if post.ID > id {
			id = post.ID
		}
	}
	return id + 1
}

// Post type for a Post Instance
type Token struct {
	ID      int64  `json:"id"`
	Token   string `json:"token"`
	Expires int64  `json:"expires"`
}

// GetID returns ID of the Post
func (p *Token) GetID() int64 { return p.ID }

// SetID sets ID of the Post
func (p *Token) SetID(id int64) { p.ID = id }

// SaveToStore saves the Post to the given dB store
func (p *Token) SaveToStore(store *store.Store) int64 {
	posts := new(Tokens)
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

func NewUUID() string {
	res := make([]byte, 20)
	rand.Read(res)
	return string(res)
}

// PostStore is database Instance of posts
var TokenStore = store.New("posts")

type server struct{}

func (s *server) CreateToken(ctx context.Context, req *authpb.CreateTokenRequest) (*authpb.CreateTokenResponse, error) {
	ts := time.Now().Unix() + 3600
	token := Token{Token: NewUUID(), Expires: int64(ts)}
	id := token.SaveToStore(TokenStore)
	created := &Tokens{}
	created.ReadFromDb()
	createdPost := created.GetByID(id)

	return &authpb.CreateTokenResponse{
		Token:     createdPost.Token,
		Expires:   createdPost.Expires,
	}, nil
}

func main() {
	log.Println("Starting auth Server")
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
	TokenStore.SetFilePath(abs)

	log.Println("Starting Posts server on port 8003")
	lis, err := net.Listen("tcp", "0.0.0.0:8003")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()

	authpb.RegisterAuthServiceServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
