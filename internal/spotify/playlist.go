package spotify

import (
	"context"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/zmb3/spotify/v2"
)

// Create
// Delete is not supported
// List
// Update
// Sort --> Update/Create new
// Filter --> update/create new
// Play?

// CreatePlaylist creates a new playlist for the user
func CreatePlaylist(client *spotify.Client, name string) error {
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		return err
	}
	_, err = client.CreatePlaylistForUser(context.Background(), user.ID, name, "", false, false)
	return err
}

// ListPlaylists lists the current user's playlists
func ListPlaylists(client *spotify.Client) ([]spotify.SimplePlaylist, error) {
	playlists, err := client.CurrentUsersPlaylists(context.Background())
	if err != nil {
		return nil, err
	}
	return playlists.Playlists, nil
}

func GetPlaylistTrackIDs(client *spotify.Client, playlistID spotify.ID) ([]spotify.ID, error) {
	// Fetch the playlist items
	playlistItems, err := client.GetPlaylistItems(context.Background(), playlistID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch playlist tracks: %w", err)
	}

	// Assuming all items are tracks, get IDs
	var trackIDs []spotify.ID
	for _, item := range playlistItems.Items {
		trackIDs = append(trackIDs, item.Track.Track.ID)
	}

	return trackIDs, nil
}


// Print will print the FullPlaylist to the given ioWriter, including the 
// song name, artist(s), and album
func PrintPlaylist(w io.Writer, pl spotify.FullPlaylist) {
	tabW :=  tabwriter.NewWriter(w, 0, 4, 2, ' ', 0)

	// Print Columns
	fmt.Fprintln(tabW, strings.Join(trackColumns, "\t"))
	
	// Print Content
	for _, tr := range(pl.Tracks.Tracks) {
		track := tr.Track
		fmt.Fprintf(tabW, "%s\t%s\t%s\n", track.Name, names(track.Artists), track.Album.Name)
	}

	tabW.Flush()
}