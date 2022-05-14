package api

import (
	"testing"
)

func TestFetchRepoDetails(t *testing.T) {
	type args struct {
		username string
		repo     string
	}
	tests := []struct {
		name string
		args args
		want RepoDetails
	}{
		{name: "Test Fetch Repo Details", args: args{username: "rupakkarki27", repo: "object_detection"}, want: RepoDetails{Name: "object_detection", FullName: "rupakkarki27/object_detection"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := FetchRepoDetails(tt.args.username, tt.args.repo)

			if got.Name != tt.want.Name || got.FullName != tt.want.FullName {
				t.Errorf("FetchRepoDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
