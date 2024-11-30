package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(trackCmd)
}

var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Manage tracks in playlists",
}
