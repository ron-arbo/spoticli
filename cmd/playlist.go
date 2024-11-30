package cmd

import (
	"fmt"

	"github.com/ron-arbo/spoticli/internal/auth"
	"github.com/ron-arbo/spoticli/internal/spotify"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(playlistCmd)
	playlistCmd.AddCommand(createPlaylistCmd)
	playlistCmd.AddCommand(listPlaylistsCmd)
}

var playlistCmd = &cobra.Command{
	Use:   "playlist",
	Short: "Perform CRUD operations on playlists",
}

// Create a new playlist
var createPlaylistCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new playlist",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		playlistName := args[0]
		client := auth.GetSpotifyClient()
		err := spotify.CreatePlaylist(client, playlistName)
		if err != nil {
			fmt.Println("Error creating playlist:", err)
		} else {
			fmt.Println("Playlist created:", playlistName)
		}
	},
}

// List playlists
var listPlaylistsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all playlists",
	Run: func(cmd *cobra.Command, args []string) {
		client := auth.GetSpotifyClient()
		playlists, err := spotify.ListPlaylists(client)
		if err != nil {
			fmt.Println("Error listing playlists:", err)
			return
		}
		for _, playlist := range playlists {
			fmt.Printf("%s (%d tracks)\n", playlist.Name, playlist.Tracks.Total)
		}
	},
}
