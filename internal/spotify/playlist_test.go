package spotify

import (
	"bytes"
	"testing"

	"github.com/zmb3/spotify/v2"
)

// Prepare mock data
var (
	mockPlaylist = spotify.FullPlaylist{
		Tracks: spotify.PlaylistTrackPage{
			Tracks: []spotify.PlaylistTrack{
				{
					Track: spotify.FullTrack{
						SimpleTrack: spotify.SimpleTrack{
							Name: "Track1",
							Artists: []spotify.SimpleArtist{
								{
									Name: "Artist1",
								},
							},
						},
						Album: spotify.SimpleAlbum{Name: "Album1"},
					},
				},
				{
					Track: spotify.FullTrack{
						SimpleTrack: spotify.SimpleTrack{
							Name: "Track2",
							Artists: []spotify.SimpleArtist{
								{
									Name: "Artist2",
								},
							},
						},
						Album: spotify.SimpleAlbum{Name: "Album2"},
					},
				},
			},
		},
	}
)

func TestPrint(t *testing.T) {
	type args struct {
		pl spotify.FullPlaylist
	}
	tests := []struct {
		name string
		args args
		expect string
	}{
		{
			name: "Test Print",
			args: args{
				mockPlaylist,
			},
			// 4 spaces between
			expect: `Song    Artist   Album
Track1  Artist1  Album1
Track2  Artist2  Album2
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect output to this buffer so we can check output
			var output bytes.Buffer
			PrintPlaylist(&output, tt.args.pl)

			if tt.expect != output.String() {
                t.Errorf("got: \n%s but expected:\n%s", output.String(), tt.expect)
            }
		})
	}
}
