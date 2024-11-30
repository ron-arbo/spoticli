package spotify

import (
	"context"
	"fmt"
	"sort"

	"github.com/zmb3/spotify/v2"
)

var (
	audioFeaturesColumns = []string{"acousticness","danceability" ,"energy","instrumentalness","liveness","loudness",
	 "speechiness", "tempo", "valence"}
)


// TODO: Does adding a type here just make things complicated?
type AudioFeature string
const (
	acousticness AudioFeature = "acousticness"
	danceability AudioFeature = "danceability"
	energy AudioFeature = "energy"
	instrumentalness AudioFeature = "instrumentalness"
	liveness AudioFeature = "liveness"
	loudness AudioFeature = "loudness"
	speechiness AudioFeature = "speechiness"
	tempo AudioFeature = "tempo"
	valence AudioFeature = "valence"
)

func GetAudioFeaturesByBatch(client *spotify.Client, trackIDs []spotify.ID) ([]*spotify.AudioFeatures, error){
	// client.GetAudioFeatures can get up to 100 IDs per call, so if our list is > 100 
	// items we should get the audio features in batches
	const batchSize = 100
	var allAudioFeatures []*spotify.AudioFeatures
	for i := 0; i < len(trackIDs); i += batchSize {
		end := i + batchSize
		if end > len(trackIDs) {
			end = len(trackIDs)
		}

		audioFeatures, err := client.GetAudioFeatures(context.Background(), trackIDs[i:end]...)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch audio features: %w", err)
		}

		allAudioFeatures = append(allAudioFeatures, audioFeatures...)
	}

	return allAudioFeatures, nil
}

// GetAudioFeaturesForPlaylist retrieves the AudioFeatures for each track in a given playlist
func GetAudioFeaturesForPlaylist(client *spotify.Client, playlistID spotify.ID) ([]*spotify.AudioFeatures, error) {
	// Fetch the playlist items
	trackIDs, err := GetPlaylistTrackIDs(client, playlistID)
	if err != nil {
		return nil, err
	}

	audioFeatures, err := GetAudioFeaturesByBatch(client, trackIDs)
	if err != nil {
		return nil, err
	}

	return audioFeatures, nil
}

// SortPlaylistByAudioFeature sorts a playlist by a specific audio feature and reorders it in Spotify
func SortPlaylistByAudioFeature(client *spotify.Client, playlistID spotify.ID, feature string, ascending bool) error {
	// Get the track IDs
	trackIDs, err := GetPlaylistTrackIDs(client, playlistID)
	if err != nil {
		return err
	}

	// Get the audio features
	audioFeatures, err := GetAudioFeaturesForPlaylist(client, playlistID)
	if err != nil {
		return err
	}

	// Connect the IDs and audioFeatures under a common object
	type trackInfo struct {
		ID        spotify.ID
		Feature   float32
	}
	var tracks []trackInfo

	// Sanity check - make sure our slice lengths line up
	if len(trackIDs) != len(audioFeatures) {
		return fmt.Errorf("trackIDs (%d) and audioFeatures (%d) slices have different lengths",
			len(trackIDs), len(audioFeatures))
	}

	// Populate a trackInfo list
	for i := range(trackIDs) {
		featureValue, err := getAudioFeatureByString(*audioFeatures[i], AudioFeature(feature))
		if err != nil {
			return err
		}

		tracks = append(tracks, trackInfo{
			ID: trackIDs[i],
			Feature: featureValue,
		})
	}


	// Sort tracks based on the feature value
	sort.Slice(tracks, func(i, j int) bool {
		if ascending {
			return tracks[i].Feature < tracks[j].Feature
		}
		return tracks[i].Feature > tracks[j].Feature
	})

	// Reorder the playlist in Spotify
	var reorderedTrackIDs []spotify.ID
	for _, track := range tracks {
		reorderedTrackIDs = append(reorderedTrackIDs, track.ID)
	}
	err = reorderPlaylist(client, playlistID, reorderedTrackIDs)
	if err != nil {
		return fmt.Errorf("failed to reorder playlist: %w", err)
	}

	fmt.Println("Playlist reordered successfully!")
	return nil
}

// reorderPlaylist reorders tracks in the playlist
func reorderPlaylist(client *spotify.Client, playlistID spotify.ID, trackIDs []spotify.ID) error {
	// Clear existing tracks
	// TODO: Maxed out at 100 tracks here
	err := client.ReplacePlaylistTracks(context.Background(), spotify.ID(playlistID), trackIDs...)
	if err != nil {
		return err
	}
	return nil
}
 
// Only support sfloat32 audiofeatures
func getAudioFeatureByString(audioFeatures spotify.AudioFeatures, audioFeature AudioFeature) (float32, error) {
	switch audioFeature {
	case acousticness:
		return audioFeatures.Acousticness, nil
	case danceability:
		return audioFeatures.Danceability, nil
	case energy:
		return audioFeatures.Energy, nil
	case instrumentalness:
		return audioFeatures.Instrumentalness, nil
	case liveness:
		return audioFeatures.Liveness, nil
	case loudness:
		return audioFeatures.Loudness, nil
	case speechiness:
		return audioFeatures.Speechiness, nil
	case tempo:
		return audioFeatures.Tempo, nil
	case valence:
		return audioFeatures.Valence, nil
	default:
		return 0, fmt.Errorf("unrecognized audio feature %s", audioFeature)
	}
}