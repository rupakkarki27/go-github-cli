package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rupakkarki27/go-github-cli/pkg/helpers"
)

type RepoDetails struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Owner    struct {
		Login   string `json:"login"`
		HTMLURL string `json:"html_url"`
	}
	HTMLURL         string    `json:"html_url"`
	OpenIssuesCount int       `json:"open_issues_count"`
	ForksCount      int       `json:"forks_count"`
	StargazersCount int       `json:"stargazers_count"`
	Language        string    `json:"language"`
	Size            int       `json:"string"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	PushedAt        time.Time `json:"pushed_at"`
}

func FetchRepoDetails(username string, repo string) RepoDetails {
	repodetails := RepoDetails{}
	url := "https://api.github.com/repos/" + username + "/" + repo

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("An error occured")
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		os.Exit(1)
	}

	defer resp.Body.Close()

	helpers.DecodeJSON(resp.Body, &repodetails)

	return repodetails
}
