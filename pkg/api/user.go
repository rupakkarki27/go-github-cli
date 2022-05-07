package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rupakkarki27/go-github-cli/pkg/helpers"
)

type UserDetails struct {
	Login        string    `json:"login"`
	ID           int       `json:"id"`
	NodeID       string    `json:"node_id"`
	AvatarURL    string    `json:"avatar_url"`
	GravatarURL  string    `json:"gravatar_url"`
	URL          string    `json:"url"`
	HTMLURL      string    `json:"html_url"`
	FollowersURL string    `json:"followers_url"`
	FollowingURL string    `json:"following_url"`
	GistsURL     string    `json:"gists_url"`
	StarredURL   string    `json:"starred_url"`
	PublicRepos  int       `json:"public_repos"`
	PublicGists  int       `json:"public_gists"`
	Followers    int       `json:"followers"`
	Following    int       `json:"following"`
	CreatedAt    time.Time `json:"created_at"`
}

func FetchUserDetails(username string) UserDetails {

	user := UserDetails{}

	url := "https://api.github.com/users/" + username

	resp, err := http.Get(url)

	if resp.StatusCode == http.StatusNotFound {
		fmt.Println("Cannot find username. Exiting program...")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("An Error occured")
	}

	defer resp.Body.Close()

	helpers.DecodeJSON(resp.Body, &user)

	return user
}
