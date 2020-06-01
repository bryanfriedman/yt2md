package youtube

import (
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/bryanfriedman/yt2mdhelpers"
)

// Video values
type Video struct {
	Title       string
	Description string
	FileName    string
	VideoID     string
	Hosts       []string
	Guests      []string
	Date        string
	PublishDate string
	Image       string
}

// WriteMarkdownFile creates a Markdown file based on the given parameters
func WriteMarkdownFile(yt YouTube, videoID string, fileName string) {

	snippet := yt.getVideoSnippet(videoID)

	var video Video
	video.VideoID = videoID
	video.FileName = fileName
	video.Title = snippet.Title
	video.Description = snippet.Description
	video.Date = snippet.PublishedAt
	video.PublishDate = time.Now().Format("2006-01-02")

	if yt.Thumbs {
		imageFilename := fileName + ".jpg"
		url := snippet.Thumbnails.Medium.Url
		err := helpers.DownloadFile(url, yt.Paths.ImagePath+imageFilename)
		helpers.HandleError(err, "Error downloading image file")
		video.Image = "/" + yt.Paths.ImagePath + imageFilename
	} else {
		video.Image = ""
	}

	t := template.Must(template.New(yt.Paths.TemplateFile).ParseFiles(yt.Paths.TemplatePath + yt.Paths.TemplateFile))
	f, err := os.Create(yt.Paths.MarkdownPath + fileName + ".md")
	helpers.HandleError(err, "Error writing Markdown file")
	defer f.Close()

	//err = t.Execute(os.Stdout, video)
	err = t.Execute(f, video)
	helpers.HandleError(err, "Error applying template")
}

// WriteMarkdownFiles creates multiple Markdown files based on the given playlist
func WriteMarkdownFiles(yt YouTube, playlistID string) {

	videoIDs := yt.getPlaylistVideoIDList(playlistID)
	for i, videoID := range videoIDs {
		episodeNumber := fmt.Sprintf("%04d", i+1)
		WriteMarkdownFile(yt, videoID, episodeNumber)
	}

}
