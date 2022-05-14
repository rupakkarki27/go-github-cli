package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rupakkarki27/go-github-cli/pkg/handlers"
)

var username = flag.String("username", " ", "Username of the GitHub user")
var show_user_repos = flag.Bool("show-user-repos", false, "Show repo details of the user")
var sort_by = flag.String("sort-by", " ", "Specify the option to sort repos by: created | updated | watchers | issues | forks")
var show_repo = flag.Bool("show-repo", false, "Show details of a specific repo")
var repo = flag.String("repo", " ", "Name of the repo whose details are needed")

func main() {
	flag.Parse()

	if *username == " " {
		fmt.Println("Error! no username specified. Use --username flag to specify username")
		os.Exit(1)
	}

	handlers.HandleUserDetails(*username)

	if *show_user_repos {
		handlers.HandleRepos(*username, *sort_by)
	}

	if *show_repo && *repo != " " {
		handlers.HandleRepoDetails(*username, *repo)
	}

}
