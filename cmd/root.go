package cmd

import (
	"fmt"
	"os"

	"github.com/bryanfriedman/yt2mdhelpers"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	homedir "github.com/mitchellh/go-homedir"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "yt2md",
	Short: "Turn a YouTube video or playlist into Markdown file(s) to add to SSG site",
	Long:  `Turn a YouTube video or playlist into Markdown file(s) to add to SSG site`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.yt2md.yaml)")

	rootCmd.PersistentFlags().StringP("apiKey", "k", "", "A YouTube API key to authenticate (i.e. AIza...)")
	rootCmd.PersistentFlags().StringP("markdownPath", "m", "markdown/", "Path to write Markdown files to.")
	rootCmd.PersistentFlags().StringP("imagePath", "i", "images/", "Path to write thumbnail image files to.")
	rootCmd.PersistentFlags().StringP("templatePath", "t", "templates/", "Path to folder where template files are.")
	rootCmd.PersistentFlags().StringP("templateFile", "f", "video.md", "File name of Go template for output.")
	rootCmd.PersistentFlags().BoolP("thumbs", "b", true, "Set whether you want to download thumbnail images or not.")
	err := viper.BindPFlags(rootCmd.PersistentFlags())
	helpers.HandleError(err, "")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		helpers.HandleError(err, "")

		// Search config in home directory with name ".yt2md" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(".yt2md")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
