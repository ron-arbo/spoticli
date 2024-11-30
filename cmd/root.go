package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "spoticli",
	Short: "",
	Long:  "",
}

// Execute is the entry point for the application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// Set up Viper to read from environment variables or config files
	viper.SetEnvPrefix("SPOTIFY")
	viper.AutomaticEnv()
}
