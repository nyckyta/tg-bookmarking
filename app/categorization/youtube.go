package categorization

import (
	"context"
	"errors"
	"net/url"
	"regexp"
	"strings"

	"google.golang.org/api/option"
	youtube "google.golang.org/api/youtube/v3"
)

type YoutubeTagsFetcher struct {
	ApiKey string
}

// fetches list of tags from Youtube video
// request must be an url to Youtube video in format https://www.youtube.com/watch?v=... or https://youtu.be/...
func (fetcher *YoutubeTagsFetcher) Fetch(request string) ([]string, error) {
	urlValidationErr := validateYoutubeUrl(request)
	if urlValidationErr != nil {
		return nil, urlValidationErr
	}

	videoId := extractVideoIdFromUrl(request)
	if videoId == "" {
		return nil, errors.New("invalid Youtube URL")
	}

	service, err := youtube.NewService(context.TODO(), option.WithAPIKey(fetcher.ApiKey))

	if err != nil {
		panic("Error creating new Youtube service")
	}

	resp, err := service.Videos.List([]string{"snippet"}).Id(videoId).Do()

	if err != nil {
		return nil, err
	}

	if len(resp.Items) == 0 {
		return nil, errors.New("video not found")
	}

	return resp.Items[0].Snippet.Tags, nil
}

func IsYotubueUrl(request string) bool {
	err := validateYoutubeUrl(request)
	return err == nil
}

// validates that string is a valid Youtube video url
func validateYoutubeUrl(request string) error {
	loweredRequest := strings.ToLower(request)
	// took from https://regexr.com/3dj5t
	reg, err := regexp.Compile(`^((?:https?:)?\/\/)?((?:www|m)\.)?((?:youtube\.com|youtu.be))(\/(?:[\w\-]+\?v=|embed\/|v\/)?)([\w\-]+)(\S+)?$`)
	
	if err != nil {
		panic("Error compiling Youtube url regexp")
	}

	if !reg.MatchString(loweredRequest) {
		return errors.New("invalid Youtube URL")
	}

	return nil
}

// parses normalized url to Youtube video and extracts video id from it
func extractVideoIdFromUrl(videoUrl string) string {
	url, e := url.Parse(videoUrl)
	if e != nil || url.Query().Get("v") == "" {
		return ""
	}

	return url.Query().Get("v")
}