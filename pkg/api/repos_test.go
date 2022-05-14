package api

import (
	"testing"
)

func TestFetchUserRepos(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want UserRepos
	}{
		{name: "Test User Repos", args: args{username: "rupakkarki27"}, want: UserRepos{
			{FullName: "rupakkarki27/object_detection"},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FetchUserRepos(tt.args.username)
			if got[0].FullName != tt.want[0].FullName {
				t.Errorf("FetchUserRepos() = %v, want %v", got, tt.want)
			}
		})
	}
}
