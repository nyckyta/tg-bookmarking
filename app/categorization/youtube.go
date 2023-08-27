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

	keywords := make([]string, len(resp.Items[0].Snippet.Tags))
	for i, tag := range resp.Items[0].Snippet.Tags {
		keywords[i] = normalizeKeyword(tag)
	}
	return keywords, nil
}

// returns true if url is youtube url to specific video
func IsYoutubeUrlToSpecificVideo(request string) bool {
	err := validateYoutubeUrl(request)
	return err == nil
}

// validates that string is a valid Youtube video url
func validateYoutubeUrl(request string) error {
	loweredRequest := strings.ToLower(request)
	// took from https://regexr.com/3dj5t with slight improvement, specifically adding youtube-nocookie.com domain
	reg, err := regexp.Compile(`^((?:https?:)?\/\/)?((?:www|m)\.)?((?:youtube\.com|youtu.be|youtube-nocookie.com))(\/(?:[\w\-]+\?v=|embed\/|v\/)?)([\w\-]+)(\S+)?$`)

	if err != nil {
		panic("Error compiling Youtube url regexp")
	}

	if !reg.MatchString(loweredRequest) {
		return errors.New("invalid Youtube URL")
	}

	return nil
}

// parses normalized url to Youtube video and extracts video id from it, it expects that URL is always valid and thus has a video id
// otherwise it will return empty string
func extractVideoIdFromUrl(videoUrl string) string {
	url, e := url.Parse(videoUrl)

	if e != nil {
		return ""
	}

	// handle the case when video id is embeded as url parameter
	if url.Query().Get("v") != "" {
		return url.Query().Get("v")
	}

	pathTokens := strings.Split(url.Path, "/")
	if len(pathTokens) > 0 && pathTokens[0] == "" {
		pathTokens = pathTokens[1:]
	}

	if len(pathTokens) == 1 {
		// weirdest from all urls, can path contain ampersands?
		// TODO: figure this format out
		if strings.HasSuffix(pathTokens[0], "&feature=channel") {
			return strings.TrimSuffix(pathTokens[0], "&feature=channel")
		}
		// handle url like {domain}/oembed?url={another youtube url with video id}
		if pathTokens[0] == "oembed" && url.Query().Has("url") {
			parameterUrl := url.Query().Get("url")
			return extractVideoIdFromUrl(parameterUrl)
		}

		// handle url like {domain}/attribution_link?...&u={another youtube url with video id}
		if pathTokens[0] == "attribution_link" && url.Query().Has("u") {
			modifiedUrl := "https://www.youtube.com" + url.Query().Get("u")
			return extractVideoIdFromUrl(modifiedUrl)
		}

		// handles the {domain}/{VIDEO_ID} case
		return pathTokens[0]
	}

	// handles the {domain}/(embed|v|e|oembed)/{VIDEO_ID} case
	if pathTokens[0] == "embed" || pathTokens[0] == "v" || pathTokens[0] == "e" {
		return pathTokens[1]
	}

	return ""
}
