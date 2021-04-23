package main_test

import (
	"testing"

	main "github.com/n3k0fi5t/travisci_demo"
)

func TestSomeWork(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "single Pass",
			args: args{1},
			want: true,
		},
		{
			name: "multiple Pass",
			args: args{5},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := main.SomeWork(tt.args.n); got != tt.want {
				t.Errorf("SomeWork() = %v, want %v", got, tt.want)
			}
		})
	}
}
