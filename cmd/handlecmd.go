package cmd

import (
	"errors"
	"os"
	"strings"

	"github.com/bryanfriedman/yt2md/helpers"
	"github.com/bryanfriedman/yt2md/youtube"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func handleCmd(cmd *cobra.Command, args []string) youtube.YouTube {

	apiKey := viper.GetString("apiKey")
	markdownPath := viper.GetString("markdownPath")
	imagePath := viper.GetString("imagePath")
	templatePath := viper.GetString("templatePath")
	templateFile := viper.GetString("templateFile")
	thumbs := viper.GetBool("thumbs")

	if apiKey == "" {
		helpers.HandleError(errors.New("Use --apiKey or define in config file"), "Must provide an API key")
	}

	if !strings.HasSuffix(markdownPath, "/") {
		markdownPath = markdownPath + "/"
	}
	if !strings.HasSuffix(imagePath, "/") {
		imagePath = imagePath + "/"
	}
	if !strings.HasSuffix(templatePath, "/") {
		templatePath = templatePath + "/"
	}

	if _, err := os.Stat(markdownPath); os.IsNotExist(err) {
		os.Mkdir(markdownPath, 0755)
	}

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		os.Mkdir(imagePath, 0755)
	}

	var yt youtube.YouTube
	yt.APIKey = apiKey
	yt.Connect()

	yt.Thumbs = thumbs
	yt.Paths.ImagePath = imagePath
	yt.Paths.MarkdownPath = markdownPath
	yt.Paths.TemplateFile = templateFile
	yt.Paths.TemplatePath = templatePath

	return yt
}
