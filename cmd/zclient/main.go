package main

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/dpakach/zwiter/client"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

var pizza = []string{"Zwitter is the best"}

func info() {
	app.Name = "Zwitter Cli Client"
	app.Usage = "Cli for accessing zwitter"
	app.Author = "dpakach"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "create-user",
			Aliases: []string{"uc"},
			Usage:   "Create a new User",
			Action: func(c *cli.Context) error {
				cc, uc := client.NewUsersClient()
				defer cc.Close()
				if c.NArg() < 1 {
					return errors.New("Not Enough arguments to create user. provide username")
				}
				user := client.CreateUser(uc, c.Args().Get(0))
				log.Printf("Response from server:\n%v\n", string(user))
				return nil
			},
		},
		{
			Name:    "get-users",
			Aliases: []string{"ug"},
			Usage:   "Get all Users",
			Action: func(c *cli.Context) error {
				cc, uc := client.NewUsersClient()
				defer cc.Close()
				users := client.GetUsers(uc)
				log.Printf("Response from server:\n%v\n", string(users))
				return nil
			},
		},
		{
			Name:    "get-user",
			Aliases: []string{"uget"},
			Usage:   "Get a User by Id",
			Action: func(c *cli.Context) error {
				cc, uc := client.NewUsersClient()
				defer cc.Close()
				if c.NArg() < 1 {
					return errors.New("Not Enough arguments to get user. provide id")
				}
				id, err := strconv.Atoi(c.Args().First())
				if err != nil {
					return errors.New("Provide valid Id to get user")
				}
				user := client.GetUser(uc, int64(id))
				log.Printf("Response from server:\n%v\n", string(user))
				return nil
			},
		},

		{
			Name:    "create-post",
			Aliases: []string{"pc"},
			Usage:   "Create a new Post",
			Action: func(c *cli.Context) error {
				cc, uc := client.NewPostsClient()
				defer cc.Close()
				if c.NArg() < 1 {
					return errors.New("Not Enough arguments to create post. provide content")
				}
				post := client.CreatePost(uc, c.Args().Get(0))
				log.Printf("Response from server:\n%v\n", string(post))
				return nil
			},
		},
		{
			Name:    "get-posts",
			Aliases: []string{"pg"},
			Usage:   "Get all Posts",
			Action: func(c *cli.Context) error {
				cc, uc := client.NewPostsClient()
				defer cc.Close()
				posts := client.GetPosts(uc)
				log.Printf("Response from server:\n%v\n", string(posts))
				return nil
			},
		},
		{
			Name:    "get-post",
			Aliases: []string{"pget"},
			Usage:   "Get a Post by Id",
			Action: func(c *cli.Context) error {
				cc, uc := client.NewPostsClient()
				defer cc.Close()
				if c.NArg() < 1 {
					return errors.New("Not Enough arguments to get post. provide id")
				}
				id, err := strconv.Atoi(c.Args().First())
				if err != nil {
					return errors.New("Provide valid Id to get post")
				}
				post := client.GetPost(uc, int64(id))
				log.Printf("Response from server:\n%v\n", string(post))
				return nil
			},
		},
	}
}

func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
