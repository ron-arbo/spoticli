package spotify

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/zmb3/spotify/v2"
)

type FeaturedTrack struct {
	simpleTrack spotify.SimpleTrack
	audioFeatures spotify.AudioFeatures
}

func NewFeaturedTrack(st spotify.SimpleTrack, af spotify.AudioFeatures) FeaturedTrack {
	return FeaturedTrack{
		simpleTrack: st,
		audioFeatures: af,
	}
}

func (ft FeaturedTrack)  Print(w io.Writer) {
	tabW :=  tabwriter.NewWriter(w, 0, 4, 2, ' ', 0)

	track := ft.simpleTrack
	features := ft.audioFeatures

	// Print Columns
	fmt.Fprintln(tabW, strings.Join(append(trackColumns, audioFeaturesColumns...), "\t"))
	
	// TODO: More elegant way to do this like above?
	// Print Content
	fmt.Fprintf(tabW, "%s\t%s\t%s\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\n", 
		track.Name, names(track.Artists), track.Album.Name,
		features.Acousticness, features.Danceability, features.Energy, features.Instrumentalness,
		features.Liveness, features.Loudness, features.Speechiness, features.Tempo, features.Valence,
	)
	tabW.Flush()
}