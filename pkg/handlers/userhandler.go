package handlers

import (
	"fmt"

	"github.com/rupakkarki27/go-github-cli/pkg/api"
)

func HandleUserDetails(username string) {
	user := api.FetchUserDetails(username)

	fmt.Printf("Hi %s!\n", user.Login)
	fmt.Printf("You joined GitHub on %v.\n", user.CreatedAt.Year())
	fmt.Printf("You've got %d public repositories.\n", user.PublicRepos)
	fmt.Printf("You have %d followers and are following %d users.\n", user.Followers, user.Following)
}
