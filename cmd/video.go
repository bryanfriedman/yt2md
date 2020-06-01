package cmd

import (
	"errors"

	"github.com/bryanfriedman/yt2md/helpers"
	"github.com/bryanfriedman/yt2md/youtube"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var videoCmd = &cobra.Command{
	Use:   "video",
	Short: "Make single Markdown file from a single video",
	Long:  `Make single Markdown file from a single video`,
	Run: func(cmd *cobra.Command, args []string) {

		yt := handleCmd(cmd, args)

		name := viper.GetString("name")
		video := viper.GetString("video")
		if video == "" {
			helpers.HandleError(errors.New("Use --video or define in config file"), "Must provide a video")
		}
		if name == "" {
			name = "0001"
		}

		youtube.WriteMarkdownFile(yt, video, name)
	},
}

func init() {

	rootCmd.AddCommand(videoCmd)

	videoCmd.Flags().StringP("video", "v", "", "The video to retrieve videos from.")
	viper.BindPFlag("video", videoCmd.Flags().Lookup("video"))

}
