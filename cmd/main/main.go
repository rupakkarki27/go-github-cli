package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rupakkarki27/go-github-cli/pkg/handlers"
)

var username = flag.String("username", " ", "Username of the GitHub user")
var show_repos = flag.Bool("show-repos", false, "Show repo details of the user")
var sort_by = flag.String("sort-by", " ", "Specify the option to sort repos by: created | updated | watchers | issues | forks")

func main() {
	flag.Parse()

	if *username == " " {
		fmt.Println("Error! no username specified. Use --username flag to specify username")
		os.Exit(1)
	}

	handlers.HandleUserDetails(*username)

	if *show_repos {
		handlers.HandleRepoDetails(*username, *sort_by)
	}
}
