# YouTube to Markdown

Turn a YouTube video or playlist into Markdown file(s) to add to an SSG site like Hugo or Jekyll.

## Using yt2md

This project is written in Go and requires the [`go` CLI](https://golang.org/doc/install) to build and run. After you've downloaded the `go` cli, you may run it using `go run main.go` followed by the command and flags as indicated below. If you wish to build into an executable, run `go build` then run `yt2md` with the command and flags. 

### Markdown Templates

This tool uses Go templates to output Markdown. You may use the `video.md` (default) or `episode.md` templates that are provided, or make one yourself and reference it. The following fields are available (all `string`):

- **Title**: Title of the video from YouTube
- **Description**: Description of the video from YouTube
- **FileName**: Name of the Markdown/image files (without .md or .jpg extension). For playlists, this name will be appended with an integer padded with 0s to four places for each individual file in the list.
- **VideoID**: The ID from the YouTube video URL
- **Date**: The Date the video was uploaded to YouTube
- **PublishDate**: The Date that you ran this command
- **Image**: The "Medium" thumbnail of the YouTube video

### Command Line Options

Below are the flags for passing in configuration. You may also use a config file to set these parameters. At a minimum, you'll need a YouTube API key (see [https://cloud.google.com/docs/authentication/api-keys](https://cloud.google.com/docs/authentication/api-keys)) and a playlist or video ID (from a YouTube URL).

```
Usage:
  yt2md [command]

Available Commands:
  help        Help about any command
  playlist    Make Markdown files from a playlist
  video       Make single Markdown file from a single video

Flags:
  -k, --apiKey string         A YouTube API key to authenticate (i.e. AIza...)
      --config string         config file (default is $HOME/.yt2md.yaml)
  -h, --help                  help for yt2md
  -i, --imagePath string      Path to write thumbnail image files to. (default "images/")
  -m, --markdownPath string   Path to write Markdown files to. (default "markdown/")
  -n, --name string           The name to use when creating the Markdown file (without the .md extension).
  -f, --templateFile string   File name of Go template for output. (default "video.md")
  -t, --templatePath string   Path to folder where template files are. (default "templates/")
  -b, --thumbs                Set whether you want to download thumbnail images or not. (default true)
```