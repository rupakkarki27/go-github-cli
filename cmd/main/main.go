package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rupakkarki27/go-github-cli/pkg/handlers"
)

var username = flag.String("username", " ", "Username of the GitHub user")

func main() {
	flag.Parse()

	if *username == " " {
		fmt.Println("Error! no username specified. Use --username flag to specify username")
		os.Exit(1)
	}

	handlers.HandleUserDetails(*username)
}
