package handlers

import (
	"fmt"
	"sort"

	"github.com/rupakkarki27/go-github-cli/pkg/api"
)

func HandleRepos(username string, sort_by string) {
	repos := api.FetchUserRepos(username)

	fmt.Printf("\n----- A summary of your repos -----\n")

	switch sort_by {
	case " ":
		sort.Slice(repos[:], func(i, j int) bool {
			return repos[i].CreatedAt.After(repos[j].CreatedAt)
		})
	case "forks":
		sort.Slice(repos[:], func(i, j int) bool {
			return repos[i].ForksCount > repos[j].ForksCount
		})
	case "watchers":
		sort.Slice(repos[:], func(i, j int) bool {
			return repos[i].WatchersCount > repos[j].WatchersCount
		})
	case "issues":
		sort.Slice(repos[:], func(i, j int) bool {
			return repos[i].OpenIssuesCount > repos[j].OpenIssuesCount
		})
	default:
		sort.Slice(repos[:], func(i, j int) bool {
			return repos[i].CreatedAt.After(repos[j].CreatedAt)
		})
	}

	for index, value := range repos[0:5] {
		if sort_by == "forks" {
			fmt.Printf("%d - %d forks - %s\n", index+1, value.ForksCount, value.FullName)
		} else {
			fmt.Printf("%d - %s\n", index+1, value.FullName)
		}
	}
}
