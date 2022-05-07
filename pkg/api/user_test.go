package api

import (
	"reflect"
	"testing"
	"time"
)

func TestFetchUserDetails(t *testing.T) {
	created_at, _ := time.Parse(time.RFC3339, "2016-12-30T02:17:05Z")

	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want UserDetails
	}{
		{"Testing rupakkarki27", args{username: "rupakkarki27"}, UserDetails{Login: "rupakkarki27", ID: 24835110, NodeID: "MDQ6VXNlcjI0ODM1MTEw", PublicRepos: 21, PublicGists: 12, Followers: 5, Following: 4, CreatedAt: created_at}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FetchUserDetails(tt.args.username); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchUserDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
