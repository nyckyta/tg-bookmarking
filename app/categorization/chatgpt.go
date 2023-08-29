package categorization

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const GPT_MSG_MAX_LEN = 2048

type ChatGptTagsFetcher struct {
	OpenAiApiKey string
}

// generated with assistance https://mholt.github.io/curl-to-go/
type gptPayload struct {
	Model            string    `json:"model"`
	Messages         []Message `json:"messages"`
	Temperature      float64   `json:"temperature"`
	MaxTokens        int       `json:"max_tokens"`
	TopP             int       `json:"top_p"`
	FrequencyPenalty int       `json:"frequency_penalty"`
	PresencePenalty  int       `json:"presence_penalty"`
}

type gptResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index        int     `json:"index"`
		Message      Message `json:"message"`
		FinishReason string  `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromtpTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// makes request to chat gpt api about text and returns list of keywords
func (fetcher *ChatGptTagsFetcher) Fetch(text string) ([]string, error) {
	var textToSend string = text

	// if url passed to fetcher, we need to extract raw text from it and ask ChatGPT to extract keywords from text,
	// unforunately ChatGPT doesn't read urls itself
	if isUrl(text) {
		rawText, err := extractRawTextFromUrl(context.TODO(), text)
		if err != nil {
			return nil, err
		}

		rawText = strings.Trim(rawText, " \n\t")
		if rawText == "" {
			log.Printf("[ERR] No raw text found by url %s", text)
			return nil, fmt.Errorf("no raw text found by url %s", text)
		}
		textToSend = rawText
	}

	if textToSend == "" {
		return nil, fmt.Errorf("text is empty")
	}

	maxLenOfText := len(textToSend)
	if maxLenOfText > GPT_MSG_MAX_LEN {
		maxLenOfText = GPT_MSG_MAX_LEN
	}

	gptPayloadBody := gptPayload{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{Role: "system", Content: "You will be provided with a block of text, and your task is to extract a list of keywords from it and provide it as comma separated values. Exclude any punctuation and special characters from words."},
			{Role: "user", Content: textToSend[:maxLenOfText]},
		},
		Temperature:      0.5,
		MaxTokens:        GPT_MSG_MAX_LEN,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}

	payloadBytes, err := json.Marshal(gptPayloadBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling gpt payload: %w", err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+fetcher.OpenAiApiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to gpt: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Printf("[ERR] gpt response status code: %d", resp.StatusCode)
		return nil, fmt.Errorf("gpt response status code: %d", resp.StatusCode)
	}

	var gptResponse gptResponse = gptResponse{}

	err = json.NewDecoder(resp.Body).Decode(&gptResponse)
	if err != nil {
		return nil, fmt.Errorf("[ERR] error decoding gpt response: %w", err)
	}

	if len(gptResponse.Choices) == 0 {
		log.Printf("[ERR] gpt response: %+v", gptResponse)
		return nil, fmt.Errorf("no choices in gpt response")
	}

	keywords := make([]string, len(gptResponse.Choices[0].Message.Content))
	for i, v := range strings.Split(gptResponse.Choices[0].Message.Content, ",") {
		keywords[i] = normalizeKeyword(v)
	}
	// we expect single response from GPT with comma separated keywords
	return keywords, nil
}
