package cmd

import (
	"errors"

	"github.com/bryanfriedman/yt2mdhelpers"
	"github.com/bryanfriedman/yt2mdyoutube"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var videoCmd = &cobra.Command{
	Use:   "video",
	Short: "Make single Markdown file from a single video",
	Long:  `Make single Markdown file from a single video`,
	Run: func(cmd *cobra.Command, args []string) {

		yt := handleCmd(cmd, args)

		video := viper.GetString("video")
		name := viper.GetString("name")
		if video == "" {
			helpers.HandleError(errors.New("Use --video or define in config file"), "Must provide a video")
		}

		youtube.WriteMarkdownFile(yt, video, name)
	},
}

func init() {

	rootCmd.AddCommand(videoCmd)

	videoCmd.Flags().StringP("video", "v", "", "The video to retrieve videos from.")
	videoCmd.Flags().StringP("name", "n", "0001", "The name to use when creating the Markdown file.")
	viper.BindPFlag("video", videoCmd.Flags().Lookup("video"))
	viper.BindPFlag("name", videoCmd.Flags().Lookup("name"))

}
