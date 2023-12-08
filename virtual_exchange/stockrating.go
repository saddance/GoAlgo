package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

const openAiApiKey = "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"
const newsApiKey = "9e98323549544e56beed1f2a5129b0a7"

type ApiResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Article struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

type Source struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func getNews(stock string) ([]Article, error) {
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&apiKey=%s", stock, newsApiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching news: %s\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %s\n", err)
		return nil, err
	}

	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Printf("Error decoding JSON: %s\n", err)
		return nil, err
	}

	return apiResponse.Articles, nil
}

func getStockRating(prompt string) (float64, error) {
	client := openai.NewClient(openAiApiKey)

	prompt = fmt.Sprintf("%s\n\nRate this news in terms of stock investment decision (0 to 100 in floating point number, respond only with number, nothing else):\n", prompt)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return -1, err
	}
	rating := resp.Choices[0].Message.Content
	rating = strings.TrimSpace(rating)
	var ratingValue float64
	fmt.Printf(prompt)
	fmt.Sscanf(rating, "%f", &ratingValue)

	return ratingValue, nil
}

func trimContent(content string, maxTokens int) string {
	runes := []rune(content)
	if len(runes) > maxTokens {
		return string(runes[:maxTokens])
	}
	return content
}

func getContentTexts(articles []Article) []string {
	var contentTexts []string
	for _, article := range articles {
		contentTexts = append(contentTexts, article.Content)
	}
	return contentTexts
}

func GetStockRatings(stock string) (float64, error) {
	news, err := getNews(stock)
	if err != nil {
		return -1, err
	}
	contentNews := trimContent(strings.Join(getContentTexts(news), "\n\nNEXT\n\n"), 4000)
	rating, err := getStockRating(contentNews)
	return rating, nil
}

func main() {
	rating, err := GetStockRatings("Nvidia")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Rating: %f\n", rating)
}
