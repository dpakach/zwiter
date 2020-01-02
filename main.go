package main

import (
	"fmt"

	"github.com/dpakach/zwiter/client"
)

func main() {
	fmt.Println("Creating a new Post")

	cc, c := client.NewUsersClient()
	defer cc.Close()
	client.CreateUser(c)
}
