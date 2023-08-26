package categorization

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/k3a/html2text"
)

// KeyWordsFetcher is an interface for fetching keywords from by a certain request
type KeyWordsFetcher interface {
	Fetch(text string) ([]string, error)
}

// extracts raw text from html by url using html2text (https://github.com/k3a/html2text)
// returns raw text and error with empty response if something went wrong
func extractRawTextFromUrl(ctx context.Context, url string) (string, error) {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return "", err
	}

	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		return "", fmt.Errorf("failed to fetch info from the page. Status code %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	rawHtml, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return html2text.HTML2TextWithOptions(string(rawHtml)), nil
}

// normalizes keyword string by trimming spaces and converting to lowercase
func normalizeKeyword(keyword string) string {
	return strings.Trim(strings.ToLower(keyword), " \n\t")
}

// checks if string is a valid url
func isUrl(text string) bool {
	obj, err := url.Parse(text)
	if err != nil {
		return false
	}

	if obj.Scheme == "" || obj.Host == "" {
		return false
	}

	return true
}