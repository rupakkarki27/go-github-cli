package api

import "time"

type UserReposResponse []struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	FullName        string    `json:"full_name"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	WatchersCount   int       `json:"watchers_count"`
	StargazersCount int       `json:"stargazers_count"`
	OpenIssues      int       `json:"open_issues"`
	ForksCount      int       `json:"forks_count"`
	Size            int       `json:"size"`
}

func FetchUserRepos(username string) {}
