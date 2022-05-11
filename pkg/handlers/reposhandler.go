package handlers

import (
	"fmt"
	"sort"

	"github.com/rupakkarki27/go-github-cli/pkg/api"
)

func HandleRepoDetails(username string, sort_by string) {
	repos := api.FetchUserRepos(username)

	switch sort_by {
	case " ":
		sort.Slice(repos[:], func(i, j int) bool {
			return repos[i].CreatedAt.After(repos[j].CreatedAt)
		})
	case "forks":
		sort.Slice(repos[:], func(i, j int) bool {
			return repos[i].ForksCount > repos[j].ForksCount
		})
	default:
		sort.Slice(repos[:], func(i, j int) bool {
			return repos[i].CreatedAt.After(repos[j].CreatedAt)
		})
	}

	fmt.Println("\n----- A summary of your repos -----")
	for index, value := range repos[0:5] {
		if sort_by == "forks" {

			fmt.Printf("%d - %d forks - %s\n", index+1, value.ForksCount, value.FullName)
		} else {
			fmt.Printf("%d - %s\n", index+1, value.FullName)
		}
	}
}
