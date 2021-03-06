package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rupakkarki27/go-github-cli/pkg/helpers"
)

type UserRepos []struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	FullName        string    `json:"full_name"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	WatchersCount   int       `json:"watchers_count"`
	StargazersCount int       `json:"stargazers_count"`
	OpenIssuesCount int       `json:"open_issues_count"`
	ForksCount      int       `json:"forks_count"`
	Size            int       `json:"size"`
}

func FetchUserRepos(username string) UserRepos {
	repos := UserRepos{}

	url := "https://api.github.com/users/" + username + "/repos?per_page=100&sort=updated"

	resp, err := http.Get(url)

	if resp.StatusCode == http.StatusNotFound {
		fmt.Println("Cannot find username. Exiting program...")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("An error occured")
		os.Exit(1)
	}

	defer resp.Body.Close()

	helpers.DecodeJSON(resp.Body, &repos)

	return repos
}
