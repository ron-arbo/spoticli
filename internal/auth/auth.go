package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

var (
	authenticator *spotifyauth.Authenticator
	ch            = make(chan *spotify.Client)
	state         = "abc123" // Random state to prevent CSRF attacks

	// Scopes defines the permissions of the Spotify client
	scopes = []string{
		spotifyauth.ScopeImageUpload,
		spotifyauth.ScopePlaylistReadPrivate,
		spotifyauth.ScopePlaylistModifyPublic,
		spotifyauth.ScopePlaylistModifyPrivate,
		spotifyauth.ScopePlaylistReadCollaborative,
		spotifyauth.ScopeUserFollowModify,
		spotifyauth.ScopeUserFollowRead,
		spotifyauth.ScopeUserLibraryModify,
		spotifyauth.ScopeUserLibraryRead ,
		spotifyauth.ScopeUserReadPrivate,
		spotifyauth.ScopeUserReadEmail,
		spotifyauth.ScopeUserReadCurrentlyPlaying,
		spotifyauth.ScopeUserReadPlaybackState,
		spotifyauth.ScopeUserModifyPlaybackState,
		spotifyauth.ScopeUserReadRecentlyPlayed,
		spotifyauth.ScopeUserTopRead,
		spotifyauth.ScopeStreaming,
	}
)

// InitAuthenticator initializes the Spotify Authenticator with required scopes
func InitAuthenticator() {
	redirectURI := os.Getenv("SPOTIFY_REDIRECT_URI")
	authenticator = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(scopes...))
}

// StartAuthServer starts an HTTP server to handle Spotify OAuth2 authentication
func StartAuthServer() {
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Redirect user to Spotify's authorization page
		url := authenticator.AuthURL(state)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	})

	// Start the server on localhost:8080
	fmt.Println("Starting server, navigate to localhost:8080 to authenticate")
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
}

// completeAuth handles the Spotify callback and authenticates the user
func completeAuth(w http.ResponseWriter, r *http.Request) {
	// Verify the state to prevent CSRF attacks
	tok, err := authenticator.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatalf("Failed to get token: %v", err)
		return
	}

	// Use the token to create a new Spotify client
	client := spotify.New(authenticator.Client(r.Context(), tok))

	// Notify the main function that we have authenticated
	fmt.Fprintf(w, "Login Completed!")
	ch <- client
}

// GetSpotifyClient waits for the user to authenticate and returns a Spotify client
func GetSpotifyClient() *spotify.Client {
	return <-ch
}
