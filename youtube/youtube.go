package youtube

import (
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"

	"github.com/bryanfriedman/yt2md/helpers"
)

// PathInfo values (with trailing slashes)
type PathInfo struct {
	MarkdownPath string
	ImagePath    string
	TemplatePath string
	TemplateFile string
}

// YouTube client
type YouTube struct {
	APIKey  string
	Paths   PathInfo
	Thumbs  bool
	service *youtube.Service
}

// Connect creates a YouTube connection
func (yt *YouTube) Connect() {
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(yt.APIKey))
	helpers.HandleError(err, "Could not connect to YouTube")
	yt.service = service
}

// Retrieve playlistItems in the specified playlist
func (yt *YouTube) getPlaylistItemsList(part string, playlistID string, pageToken string) *youtube.PlaylistItemListResponse {
	call := yt.service.PlaylistItems.List(part)
	call = call.PlaylistId(playlistID)
	if pageToken != "" {
		call = call.PageToken(pageToken)
	}
	response, err := call.Do()
	helpers.HandleError(err, "Error getting playlist items")
	return response
}

func (yt *YouTube) getPlaylistVideoIDList(playlistID string) []string {
	idList := []string{}
	nextPageToken := ""
	for {

		playlistResponse := yt.getPlaylistItemsList("snippet", playlistID, nextPageToken)

		for _, playlistItem := range playlistResponse.Items {
			idList = append(idList, playlistItem.Snippet.ResourceId.VideoId)
		}

		// Set the token to retrieve the next page of results
		// or exit the loop if all results have been retrieved.
		nextPageToken = playlistResponse.NextPageToken
		if nextPageToken == "" {
			break
		}

	}

	return idList

}

// Retrieve the given video by ID
func (yt *YouTube) getVideoSnippet(videoID string) *youtube.VideoSnippet {
	call := yt.service.Videos.List("snippet").Id(videoID)
	response, err := call.Do()
	helpers.HandleError(err, "Error getting video")
	return response.Items[0].Snippet
}
