package cmd

import (
	"errors"

	"github.com/bryanfriedman/yt2mdhelpers"
	"github.com/bryanfriedman/yt2mdyoutube"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var playlistCmd = &cobra.Command{
	Use:   "playlist",
	Short: "Make Markdown files from a playlist",
	Long:  `Make Markdown files from a playlist`,
	Run: func(cmd *cobra.Command, args []string) {

		yt := handleCmd(cmd, args)

		playlist := viper.GetString("playlist")
		if playlist == "" {
			helpers.HandleError(errors.New("Use --playlist or define in config file"), "Must provide a playlist")
		}

		youtube.WriteMarkdownFiles(yt, playlist)

	},
}

func init() {
	rootCmd.AddCommand(playlistCmd)

	playlistCmd.Flags().StringP("playlist", "p", "", "The playlist to retrieve videos from.")
	viper.BindPFlag("playlist", playlistCmd.Flags().Lookup("playlist"))

}
