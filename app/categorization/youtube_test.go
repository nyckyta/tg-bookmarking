package categorization

import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

// test sets are taken from https://gist.github.com/rodrigoborgesdeoliveira/987683cfbfcc8d800192da1e73adc486
// each set is divided by video id i.e. each set contains examples with the same video id

// videoId = -wtIMTCHWuI
var youtubeUrls_set1 = []string{
	"http://www.youtube.com/watch?v=-wtIMTCHWuI",
	"http://youtube.com/watch?v=-wtIMTCHWuI",
	"http://m.youtube.com/watch?v=-wtIMTCHWuI",
	"http://www.youtube.com/v/-wtIMTCHWuI?version=3&autohide=1",
	"http://youtube.com/v/-wtIMTCHWuI?version=3&autohide=1",
	"http://m.youtube.com/v/-wtIMTCHWuI?version=3&autohide=1",
	"https://www.youtube.com/v/-wtIMTCHWuI?version=3&autohide=1",
	"https://youtube.com/v/-wtIMTCHWuI?version=3&autohide=1",
	"https://m.youtube.com/v/-wtIMTCHWuI?version=3&autohide=1",
	"http://youtu.be/-wtIMTCHWuI",
	"https://youtu.be/-wtIMTCHWuI",
	"http://www.youtube.com/oembed?url=http%3A//www.youtube.com/watch?v%3D-wtIMTCHWuI&format=json",
	"http://youtube.com/oembed?url=http%3A//www.youtube.com/watch?v%3D-wtIMTCHWuI&format=json",
	"http://m.youtube.com/oembed?url=http%3A//www.youtube.com/watch?v%3D-wtIMTCHWuI&format=json",
	"https://www.youtube.com/oembed?url=http%3A//www.youtube.com/watch?v%3D-wtIMTCHWuI&format=json",
	"https://youtube.com/oembed?url=http%3A//www.youtube.com/watch?v%3D-wtIMTCHWuI&format=json",
	"https://m.youtube.com/oembed?url=http%3A//www.youtube.com/watch?v%3D-wtIMTCHWuI&format=json",
}

// videoId = lalOy8Mbfdc
var youtubeUrls_set2 = []string{
	"https://www.youtube.com/watch?v=lalOy8Mbfdc",
	"https://youtube.com/watch?v=lalOy8Mbfdc",
	"https://m.youtube.com/watch?v=lalOy8Mbfdc",
	"http://www.youtube.com/watch?v=lalOy8Mbfdc&playnext_from=TL&videos=osPknwzXEas&feature=sub",
	"http://youtube.com/watch?v=lalOy8Mbfdc&playnext_from=TL&videos=osPknwzXEas&feature=sub",
	"http://m.youtube.com/watch?v=lalOy8Mbfdc&playnext_from=TL&videos=osPknwzXEas&feature=sub",
	"https://www.youtube.com/watch?v=lalOy8Mbfdc&playnext_from=TL&videos=osPknwzXEas&feature=sub",
	"https://youtube.com/watch?v=lalOy8Mbfdc&playnext_from=TL&videos=osPknwzXEas&feature=sub",
	"https://m.youtube.com/watch?v=lalOy8Mbfdc&playnext_from=TL&videos=osPknwzXEas&feature=sub",
	"http://www.youtube.com/watch?v=lalOy8Mbfdc&feature=youtu.be",
	"http://youtube.com/watch?v=lalOy8Mbfdc&feature=youtu.be",
	"http://m.youtube.com/watch?v=lalOy8Mbfdc&feature=youtu.be",
	"https://www.youtube.com/watch?v=lalOy8Mbfdc&feature=youtu.be",
	"https://youtube.com/watch?v=lalOy8Mbfdc&feature=youtu.be",
	"https://m.youtube.com/watch?v=lalOy8Mbfdc&feature=youtu.be",
	"http://youtu.be/lalOy8Mbfdc?t=1",
	"http://youtu.be/lalOy8Mbfdc?t=1s",
	"https://youtu.be/lalOy8Mbfdc?t=1",
	"https://youtu.be/lalOy8Mbfdc?t=1s",
	"http://www.youtube.com/embed/lalOy8Mbfdc",
	"http://youtube.com/embed/lalOy8Mbfdc",
	"http://m.youtube.com/embed/lalOy8Mbfdc",
	"https://www.youtube.com/embed/lalOy8Mbfdc",
	"https://youtube.com/embed/lalOy8Mbfdc",
	"https://m.youtube.com/embed/lalOy8Mbfdc",
	"http://www.youtube-nocookie.com/embed/lalOy8Mbfdc?rel=0",
	"https://www.youtube-nocookie.com/embed/lalOy8Mbfdc?rel=0",
}

// videoId = yZv2daTWRZU
var youtubeUrls_set3 = []string{
	"http://www.youtube.com/watch?v=yZv2daTWRZU&feature=em-uploademail",
	"http://youtube.com/watch?v=yZv2daTWRZU&feature=em-uploademail",
	"http://m.youtube.com/watch?v=yZv2daTWRZU&feature=em-uploademail",
	"https://www.youtube.com/watch?v=yZv2daTWRZU&feature=em-uploademail",
	"https://youtube.com/watch?v=yZv2daTWRZU&feature=em-uploademail",
	"https://m.youtube.com/watch?v=yZv2daTWRZU&feature=em-uploademail",
	"http://www.youtube.com/attribution_link?a=8g8kPrPIi-ecwIsS&u=/watch%3Fv%3DyZv2daTWRZU%26feature%3Dem-uploademail",
	"http://youtube.com/attribution_link?a=8g8kPrPIi-ecwIsS&u=/watch%3Fv%3DyZv2daTWRZU%26feature%3Dem-uploademail",
	"http://m.youtube.com/attribution_link?a=8g8kPrPIi-ecwIsS&u=/watch%3Fv%3DyZv2daTWRZU%26feature%3Dem-uploademail",
	"https://www.youtube.com/attribution_link?a=8g8kPrPIi-ecwIsS&u=/watch%3Fv%3DyZv2daTWRZU%26feature%3Dem-uploademail",
	"https://youtube.com/attribution_link?a=8g8kPrPIi-ecwIsS&u=/watch%3Fv%3DyZv2daTWRZU%26feature%3Dem-uploademail",
	"https://m.youtube.com/attribution_link?a=8g8kPrPIi-ecwIsS&u=/watch%3Fv%3DyZv2daTWRZU%26feature%3Dem-uploademail",
}

// videoId = 0zM3nApSvMg
var youtubeUrls_set4 = []string{
	"http://www.youtube.com/watch?v=0zM3nApSvMg&feature=feedrec_grec_index",
	"http://youtube.com/watch?v=0zM3nApSvMg&feature=feedrec_grec_index",
	"http://m.youtube.com/watch?v=0zM3nApSvMg&feature=feedrec_grec_index",
	"https://www.youtube.com/watch?v=0zM3nApSvMg&feature=feedrec_grec_index",
	"https://youtube.com/watch?v=0zM3nApSvMg&feature=feedrec_grec_index",
	"https://m.youtube.com/watch?v=0zM3nApSvMg&feature=feedrec_grec_index",
	"http://www.youtube.com/watch?v=0zM3nApSvMg#t=0m10s",
	"http://youtube.com/watch?v=0zM3nApSvMg#t=0m10s",
	"http://m.youtube.com/watch?v=0zM3nApSvMg#t=0m10s",
	"https://www.youtube.com/watch?v=0zM3nApSvMg#t=0m10s",
	"https://youtube.com/watch?v=0zM3nApSvMg#t=0m10s",
	"https://m.youtube.com/watch?v=0zM3nApSvMg#t=0m10s",
	"http://www.youtube.com/v/0zM3nApSvMg?fs=1&hl=en_US&rel=0",
	"http://youtube.com/v/0zM3nApSvMg?fs=1&hl=en_US&rel=0",
	"http://m.youtube.com/v/0zM3nApSvMg?fs=1&hl=en_US&rel=0",
	"https://www.youtube.com/v/0zM3nApSvMg?fs=1&amp;hl=en_US&amp;rel=0",
	"https://www.youtube.com/v/0zM3nApSvMg?fs=1&hl=en_US&rel=0",
	"https://youtube.com/v/0zM3nApSvMg?fs=1&hl=en_US&rel=0",
	"https://m.youtube.com/v/0zM3nApSvMg?fs=1&hl=en_US&rel=0",
}

// videoId = cKZDdG9FTKY
var youtubeUrls_set5 = []string{
	"http://www.youtube.com/watch?v=cKZDdG9FTKY&feature=channel",
	"http://youtube.com/watch?v=cKZDdG9FTKY&feature=channel",
	"http://m.youtube.com/watch?v=cKZDdG9FTKY&feature=channel",
}

// videoId = oTJRivZTMLs
var youtubeUrls_set6 = []string{
	"https://www.youtube.com/watch?v=oTJRivZTMLs&feature=channel",
	"https://youtube.com/watch?v=oTJRivZTMLs&feature=channel",
	"https://m.youtube.com/watch?v=oTJRivZTMLs&feature=channel",
	"http://youtu.be/oTJRivZTMLs?list=PLToa5JuFMsXTNkrLJbRlB--76IAOjRM9b",
	"https://youtu.be/oTJRivZTMLs?list=PLToa5JuFMsXTNkrLJbRlB--76IAOjRM9b",
	"http://youtu.be/oTJRivZTMLs&feature=channel",
	"https://youtu.be/oTJRivZTMLs&feature=channel",
}

// videoId = dQw4w9WgXcQ
var youtubeUrls_set7 = []string{
	"http://www.youtube.com/watch?v=dQw4w9WgXcQ&feature=youtube_gdata_player",
	"http://youtube.com/watch?v=dQw4w9WgXcQ&feature=youtube_gdata_player",
	"http://m.youtube.com/watch?v=dQw4w9WgXcQ&feature=youtube_gdata_player",
	"https://www.youtube.com/watch?v=dQw4w9WgXcQ&feature=youtube_gdata_player",
	"https://youtube.com/watch?v=dQw4w9WgXcQ&feature=youtube_gdata_player",
	"https://m.youtube.com/watch?v=dQw4w9WgXcQ&feature=youtube_gdata_player",
	"http://www.youtube.com/watch?feature=player_embedded&v=dQw4w9WgXcQ",
	"http://youtube.com/watch?feature=player_embedded&v=dQw4w9WgXcQ",
	"http://m.youtube.com/watch?feature=player_embedded&v=dQw4w9WgXcQ",
	"https://www.youtube.com/watch?feature=player_embedded&v=dQw4w9WgXcQ",
	"https://youtube.com/watch?feature=player_embedded&v=dQw4w9WgXcQ",
	"https://m.youtube.com/watch?feature=player_embedded&v=dQw4w9WgXcQ",
	"http://www.youtube.com/watch?app=desktop&v=dQw4w9WgXcQ",
	"http://youtube.com/watch?app=desktop&v=dQw4w9WgXcQ",
	"http://m.youtube.com/watch?app=desktop&v=dQw4w9WgXcQ",
	"https://www.youtube.com/watch?app=desktop&v=dQw4w9WgXcQ",
	"https://youtube.com/watch?app=desktop&v=dQw4w9WgXcQ",
	"https://m.youtube.com/watch?app=desktop&v=dQw4w9WgXcQ",
	"http://www.youtube.com/v/dQw4w9WgXcQ",
	"http://youtube.com/v/dQw4w9WgXcQ",
	"http://m.youtube.com/v/dQw4w9WgXcQ",
	"https://www.youtube.com/v/dQw4w9WgXcQ",
	"https://youtube.com/v/dQw4w9WgXcQ",
	"https://m.youtube.com/v/dQw4w9WgXcQ",
	"http://www.youtube.com/v/dQw4w9WgXcQ?feature=youtube_gdata_player",
	"http://youtube.com/v/dQw4w9WgXcQ?feature=youtube_gdata_player",
	"http://m.youtube.com/v/dQw4w9WgXcQ?feature=youtube_gdata_player",
	"https://www.youtube.com/v/dQw4w9WgXcQ?feature=youtube_gdata_player",
	"https://youtube.com/v/dQw4w9WgXcQ?feature=youtube_gdata_player",
	"https://m.youtube.com/v/dQw4w9WgXcQ?feature=youtube_gdata_player",
	"http://youtu.be/dQw4w9WgXcQ?feature=youtube_gdata_player",
	"https://youtu.be/dQw4w9WgXcQ?feature=youtube_gdata_player",
	"http://www.youtube.com/e/dQw4w9WgXcQ",
	"http://youtube.com/e/dQw4w9WgXcQ",
	"http://m.youtube.com/e/dQw4w9WgXcQ",
	"https://www.youtube.com/e/dQw4w9WgXcQ",
	"https://youtube.com/e/dQw4w9WgXcQ",
	"https://m.youtube.com/e/dQw4w9WgXcQ",
}

// videoId = ishbTyLs6ps
var youtubeUrls_set8 = []string{
	"http://www.youtube.com/watch?v=ishbTyLs6ps&list=PLGup6kBfcU7Le5laEaCLgTKtlDcxMqGxZ&index=106&shuffle=2655",
	"http://youtube.com/watch?v=ishbTyLs6ps&list=PLGup6kBfcU7Le5laEaCLgTKtlDcxMqGxZ&index=106&shuffle=2655",
	"http://m.youtube.com/watch?v=ishbTyLs6ps&list=PLGup6kBfcU7Le5laEaCLgTKtlDcxMqGxZ&index=106&shuffle=2655",
	"https://www.youtube.com/watch?v=ishbTyLs6ps&list=PLGup6kBfcU7Le5laEaCLgTKtlDcxMqGxZ&index=106&shuffle=2655",
	"https://youtube.com/watch?v=ishbTyLs6ps&list=PLGup6kBfcU7Le5laEaCLgTKtlDcxMqGxZ&index=106&shuffle=2655",
	"https://m.youtube.com/watch?v=ishbTyLs6ps&list=PLGup6kBfcU7Le5laEaCLgTKtlDcxMqGxZ&index=106&shuffle=2655",
}

// videoId = 3DEhxJLojIE_o
var youtubeUrls_set9 = []string{
	"http://www.youtube.com/attribution_link?a=JdfC0C9V6ZI&u=%2Fwatch%3Fv%3DEhxJLojIE_o%26feature%3Dshare",
	"http://youtube.com/attribution_link?a=JdfC0C9V6ZI&u=%2Fwatch%3Fv%3DEhxJLojIE_o%26feature%3Dshare",
	"http://m.youtube.com/attribution_link?a=JdfC0C9V6ZI&u=%2Fwatch%3Fv%3DEhxJLojIE_o%26feature%3Dshare",
	"https://www.youtube.com/attribution_link?a=JdfC0C9V6ZI&u=%2Fwatch%3Fv%3DEhxJLojIE_o%26feature%3Dshare",
	"https://youtube.com/attribution_link?a=JdfC0C9V6ZI&u=%2Fwatch%3Fv%3DEhxJLojIE_o%26feature%3Dshare",
	"https://m.youtube.com/attribution_link?a=JdfC0C9V6ZI&u=%2Fwatch%3Fv%3DEhxJLojIE_o%26feature%3Dshare",
}

// videoId = nas1rJpm7wY
var youtubeUrls_set10 = []string{
	"http://www.youtube.com/embed/nas1rJpm7wY?rel=0",
	"http://youtube.com/embed/nas1rJpm7wY?rel=0",
	"http://m.youtube.com/embed/nas1rJpm7wY?rel=0",
	"https://www.youtube.com/embed/nas1rJpm7wY?rel=0",
	"https://youtube.com/embed/nas1rJpm7wY?rel=0",
	"https://m.youtube.com/embed/nas1rJpm7wY?rel=0",
}

func TestIsYoutubeUrl_success(t *testing.T) {
	for _, url := range youtubeUrls_set1 {
		assert.True(t, IsYoutubeUrlToSpecificVideo(url), "Url %s is expected to be Youtube url", url)
	}

	for _, url := range youtubeUrls_set2 {
		assert.True(t, IsYoutubeUrlToSpecificVideo(url), "Url %s is expected to be Youtube url", url)
	}

	for _, url := range youtubeUrls_set3 {
		assert.True(t, IsYoutubeUrlToSpecificVideo(url), "Url %s is expected to be Youtube url", url)
	}

	for _, url := range youtubeUrls_set4 {
		assert.True(t, IsYoutubeUrlToSpecificVideo(url), "Url %s is expected to be Youtube url", url)
	}

	for _, url := range youtubeUrls_set5 {
		assert.True(t, IsYoutubeUrlToSpecificVideo(url), "Url %s is expected to be Youtube url", url)
	}

	for _, url := range youtubeUrls_set6 {
		assert.True(t, IsYoutubeUrlToSpecificVideo(url), "Url %s is expected to be Youtube url", url)
	}

	for _, url := range youtubeUrls_set7 {
		assert.True(t, IsYoutubeUrlToSpecificVideo(url), "Url %s is expected to be Youtube url", url)
	}

	for _, url := range youtubeUrls_set8 {
		assert.True(t, IsYoutubeUrlToSpecificVideo(url), "Url %s is expected to be Youtube url", url)
	}

	for _, url := range youtubeUrls_set9 {
		assert.True(t, IsYoutubeUrlToSpecificVideo(url), "Url %s is expected to be Youtube url", url)
	}

	for _, url := range youtubeUrls_set10 {
		assert.True(t, IsYoutubeUrlToSpecificVideo(url), "Url %s is expected to be Youtube url", url)
	}
}

func TestExctractVideoId_success(t *testing.T) {
	for _, url := range youtubeUrls_set1 {
		assert.Equal(t, "-wtIMTCHWuI", extractVideoIdFromUrl(url), "Video id must be euqal to -wtIMTCHWuI, url = %s", url)
	}

	for _, url := range youtubeUrls_set2 {
		assert.Equal(t, "lalOy8Mbfdc", extractVideoIdFromUrl(url), "Video id must be euqal to lalOy8Mbfdc, url = %s", url)
	}

	for _, url := range youtubeUrls_set3 {
		assert.Equal(t, "yZv2daTWRZU", extractVideoIdFromUrl(url), "Video id must be euqal to yZv2daTWRZU, url  = %s", url)
	}

	for _, url := range youtubeUrls_set4 {
		assert.Equal(t, "0zM3nApSvMg", extractVideoIdFromUrl(url), "Video id must be euqal to 0zM3nApSvMg, url = %s", url)
	}

	for _, url := range youtubeUrls_set5 {
		assert.Equal(t, "cKZDdG9FTKY", extractVideoIdFromUrl(url), "Video id must be euqal to cKZDdG9FTKY, url = %s", url)
	}

	for _, url := range youtubeUrls_set6 {
		assert.Equal(t, "oTJRivZTMLs", extractVideoIdFromUrl(url), "Video id must be euqal to oTJRivZTMLs, url = %s", url)
	}

	for _, url := range youtubeUrls_set7 {
		assert.Equal(t, "dQw4w9WgXcQ", extractVideoIdFromUrl(url), "Video id must be euqal to dQw4w9WgXcQ, url = %s", url)
	}

	for _, url := range youtubeUrls_set8 {
		assert.Equal(t, "ishbTyLs6ps", extractVideoIdFromUrl(url), "Video id must be euqal to ishbTyLs6ps, url = %s", url)
	}

	for _, url := range youtubeUrls_set9 {
		assert.Equal(t, "EhxJLojIE_o", extractVideoIdFromUrl(url), "Video id must be euqal to EhxJLojIE_o, url = %s", url)
	}

	for _, url := range youtubeUrls_set10 {
		assert.Equal(t, "nas1rJpm7wY", extractVideoIdFromUrl(url), "Video id must be euqal to nas1rJpm7wY, url = %s", url)
	}
}

func TestIsYoutubeUrl_failed(t *testing.T) {
	assert.False(t, IsYoutubeUrlToSpecificVideo("http://google.com"), "Url http://google.con is expected to be invalid Youtube url")
	assert.False(t, IsYoutubeUrlToSpecificVideo("youtub.ua"), "Url youtub.ua is expected to be invalid Youtube url")
	assert.False(t, IsYoutubeUrlToSpecificVideo("randomstring"), "Url random string is expected to be invalid Youtube url")
	assert.False(t, IsYoutubeUrlToSpecificVideo("youtube.com/"), "Url youtube.com is expected to be invalid Youtube url without video id")
	assert.False(t, IsYoutubeUrlToSpecificVideo("https://www.youtube.com/@dreamtheaterofficial"), "Url youtube.com is expected to be invalid Youtube url, since this is link to the channel")
}
