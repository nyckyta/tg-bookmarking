package main

import (
	"log"
	"os"
	"strings"
	"time"

	category "github.com/nyckyta/tg-bookmarking/app/categorization"
	tele "gopkg.in/telebot.v3"
	xurls "mvdan.cc/xurls/v2"
)

func main() {
	// init youtube service
	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == ""  {
		panic("GOOGLE_API_KEY environment variable is not set")
	}

	fetcher := category.YoutubeTagsFetcher{ApiKey: apiKey}

	// init bot
	pref := tele.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tele.OnText, func(c tele.Context) error {
		urls := xurls.Relaxed().FindAllString(c.Text(), -1)
		log.Printf("Urls found: %s\n", urls)
		if len(urls) == 0 {
			return nil	
		}

		keywords := []string{}
		for _, url := range urls {
			if !strings.HasPrefix(url, "https://") {
				url = "https://" + url
			}

			// handles only first youtube url? Will you ever paste two urls in one message?
			if category.IsYotubueUrl(url) {
				log.Printf("Youtube url found %s", url)
				urlKeyWords, err := fetcher.Fetch(url)
				if err != nil {
					log.Println(err)
					continue
				}
				keywords = append(keywords, urlKeyWords...)
			}
		}

		if len(keywords) > 0 {
			editedMsg := c.Message().Text + "\n\nKeywords:\n" + strings.Join(keywords, ", ")
			err = c.Send(editedMsg)
			if err != nil {
				log.Println(err)
			}
			c.Delete()
		}

		return nil
	});

	log.Println("Bot started")
	b.Start()
}