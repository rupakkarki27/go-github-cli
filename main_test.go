package main

import "testing"

func TestAddNumber(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2+2=4", args{2, 2}, 4},
		{"2+3=5", args{2, 3}, 5},
		{"3+3=6", args{2, 3}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddNumber(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AddNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
