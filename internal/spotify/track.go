package spotify

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/zmb3/spotify/v2"
)

var (
	trackColumns = []string{"Song", "Artist", "Album"}
)

// Print will print the FullTrack to the given ioWriter, including the
// song name, artist(s), and album
func PrintTrack(w io.Writer, track spotify.FullTrack) {
	tabW :=  tabwriter.NewWriter(w, 0, 4, 2, ' ', 0)

	// Print Columns
	fmt.Fprintln(tabW, "Song\tArtist\tAlbum")
	
	// Print Content
	fmt.Fprintf(tabW, "%s\t%s\t%s\n", track.Name, names(track.Artists), track.Album.Name)

	tabW.Flush()
}
