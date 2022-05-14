package handlers

import (
	"fmt"

	"github.com/rupakkarki27/go-github-cli/pkg/api"
)

func HandleRepoDetails(username string, repo string) {
	repodetails := api.FetchRepoDetails(username, repo)

	fmt.Printf("----- A summary of %s -----\n", repodetails.Name)
	fmt.Printf("Full Name: %s\n", repodetails.FullName)
	fmt.Printf("Repo URL: %s\n", repodetails.HTMLURL)
	fmt.Printf("Owner: %s\n", repodetails.Owner.Login)
	fmt.Printf("Language: %s\n", repodetails.Language)
	fmt.Printf("Open Issues: %d\n", repodetails.OpenIssuesCount)
	fmt.Printf("No. of Forks: %d\n", repodetails.ForksCount)
	fmt.Printf("Created: %v\n", repodetails.CreatedAt)
	fmt.Printf("Last Pushed: %v\n", repodetails.PushedAt)
}
