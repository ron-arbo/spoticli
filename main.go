package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ron-arbo/spoticli/internal/auth"
	// TODO: Better package name?
	spot "github.com/ron-arbo/spoticli/internal/spotify"
)

func main() {
	// Initialize the authenticator
	auth.InitAuthenticator()

	// Start the auth server to handle Spotify authentication
	auth.StartAuthServer()

	// Wait for the user to authenticate and get the Spotify client
	client := auth.GetSpotifyClient()

	playlists, err := spot.ListPlaylists(client)
	if err != nil {
		fmt.Println(err)
	}

	fullPlaylist, err := client.GetPlaylist(context.Background(), playlists[0].ID)
	if err != nil {
		fmt.Println(err)
	}	

	track := fullPlaylist.Tracks.Tracks[0].Track
	af, err := client.GetAudioFeatures(context.Background(), track.ID)
	if err != nil {
		panic(err)
	}
	featuredTrack := spot.NewFeaturedTrack(track.SimpleTrack, *af[0])

	fmt.Println(track.SimpleTrack.Album.Name)
	featuredTrack.Print(os.Stdout)

	

	// playlist1 := playlists[0]
	// fmt.Printf("Found playlist %s with ID %s\n", playlist1.Name, playlist1.ID)
	// fmt.Println("Sorting it")

	// audioFeatures, err := spot.GetAudioFeaturesForPlaylist(client, playlist1.ID)
	// if err != nil {
	// 	fmt.Println(err)
	// }



	// err = spot.SortPlaylistByAudioFeature(client, playlist1.ID, "energy", true)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// newPlaylist, err := client.GetPlaylist(context.TODO(), playlist1.ID)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// spot.PrintPlaylist(os.Stdout, *newPlaylist)

}