package spotify

import (
	"strings"

	"github.com/zmb3/spotify/v2"
)

// names prints the artists names in a comma separated list
func names(artists []spotify.SimpleArtist) string {
	var names []string
	for _, artist := range(artists) {
		names = append(names, artist.Name)
	}

	return strings.Join(names, ", ")
}