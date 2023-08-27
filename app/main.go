package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/goware/urlx"
	category "github.com/nyckyta/tg-bookmarking/app/categorization"
	tele "gopkg.in/telebot.v3"
	xurls "mvdan.cc/xurls/v2"
)

func main() {
	// init youtube service
	googleApiKey := os.Getenv("GOOGLE_API_KEY")
	if googleApiKey == "" {
		panic("GOOGLE_API_KEY environment variable is not set")
	}

	openaiApiKey := os.Getenv("OPENAI_API_KEY")
	if openaiApiKey == "" {
		panic("OPENAI_API_KEY environment variable is not set")
	}

	youtubeFetcher := category.YoutubeTagsFetcher{ApiKey: googleApiKey}
	gptFetcher := category.ChatGptTagsFetcher{OpenAiApiKey: openaiApiKey}

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
		log.Printf("[INFO] Received msg from %s\n", c.Sender().Username)

		// process urls initially
		urls := xurls.Relaxed().FindAllString(c.Text(), -1)
		log.Printf("[INFO] Found %d urls in message", len(urls))

		keywordsMap := map[string]bool{}
		for _, url := range urls {
			normalizedUrl, err := urlx.NormalizeString(url)
			if err != nil {
				log.Printf("[ERR] Failed to normalize url %v", err)
				return nil
			}

			if category.IsYoutubeUrlToSpecificVideo(normalizedUrl) {
				urlKeyWords, err := youtubeFetcher.Fetch(normalizedUrl)
				if err != nil {
					log.Printf("[ERR] Error fetching keywordsMap from url %s: %s", url, err)
					continue
				}

				for _, keyword := range urlKeyWords {
					keywordsMap[keyword] = true
				}
				continue
			}

			urlKeyWords, err := gptFetcher.Fetch(normalizedUrl)
			if err != nil {
				log.Printf("[ERR] Error fetching keywordsMap from url %s: %s", url, err)
				continue
			}
			for _, keyword := range urlKeyWords {
				keywordsMap[keyword] = true
			}
		}

		log.Printf("[INFO] Processed urls, found keywords %d", len(keywordsMap))

		keywords := make([]string, 0, len(keywordsMap))
		for k := range keywordsMap {
			if strings.Trim(k, " \n\t") == "" {
				continue
			}
			keywords = append(keywords, k)
		}

		log.Printf("[INFO] Found %d keywords", len(keywords))

		if len(keywordsMap) == 0 {
			return nil
		}

		editedMsg := c.Message().Text + "\n\nKeywords:\n" + strings.Join(keywords, ", ")
		err = c.Send(editedMsg)
		if err != nil {
			log.Printf("[ERR] Failed to send message %s", err)
		}

		err := c.Delete()
		if err != nil {
			log.Printf("[ERR] Failed to delete message %s", err)
		}

		return nil
	})

	log.Println("[INFO] Bot started")
	b.Start()
}
