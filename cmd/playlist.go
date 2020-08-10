package cmd

import (
	"errors"

	"github.com/bryanfriedman/yt2md/helpers"
	"github.com/bryanfriedman/yt2md/youtube"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var playlistCmd = &cobra.Command{
	Use:   "playlist",
	Short: "Make Markdown files from a playlist",
	Long:  `Make Markdown files from a playlist`,
	Run: func(cmd *cobra.Command, args []string) {

		yt := handleCmd(cmd, args)

		name := viper.GetString("name")
		playlist := viper.GetString("playlist")
		reverseSort := viper.GetBool("reverseSort")
		if playlist == "" {
			helpers.HandleError(errors.New("Use --playlist or define in config file"), "Must provide a playlist")
		}

		youtube.WriteMarkdownFiles(yt, playlist, name, reverseSort)

	},
}

func init() {
	rootCmd.AddCommand(playlistCmd)

	playlistCmd.Flags().StringP("playlist", "p", "", "The playlist to retrieve videos from.")
	playlistCmd.Flags().BoolP("reverseSort", "r", false, "Set whether to sort playlist in reverse order (oldest first).")
	viper.BindPFlag("playlist", playlistCmd.Flags().Lookup("playlist"))
	viper.BindPFlag("reverseSort", playlistCmd.Flags().Lookup("reverseSort"))
}
