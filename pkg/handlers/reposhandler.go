package handlers

import (
	"fmt"
	"sort"

	"github.com/rupakkarki27/go-github-cli/pkg/api"
)

func HandleRepoDetails(username string) {
	repos := api.FetchUserRepos(username)

	sort.Slice(repos, func(i, j int) bool {
		return repos[i].CreatedAt.After(repos[j].CreatedAt)
	})

	fmt.Println("Here are your latest 5 repos:")
	for index, value := range repos[1:6] {
		fmt.Printf("%d - %s\n", index+1, value.FullName)
	}
}
