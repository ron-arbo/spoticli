package spotify

import (
	"testing"

	"github.com/zmb3/spotify/v2"
)

func Test_names(t *testing.T) {
	type args struct {
		artists []spotify.SimpleArtist
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := names(tt.args.artists); got != tt.want {
				t.Errorf("names() = %v, want %v", got, tt.want)
			}
		})
	}
}
