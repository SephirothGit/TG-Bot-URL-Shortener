package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

type createResponse struct {
	Data struct {
		ShortCode string `json:"short_code"`
	} `json:"data"`
}

type linksResponse struct {
	Data []struct {
		OriginalURL string `json:"original_url"`
		ShortCode   string `json:"short_code"`
		Clicks      int    `json:"clicks"`
	} `json:"data"`
}

func main() {
	godotenv.Load()

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Hello! I am a URL shortener, here you can paste the link and receive a short version of it")
	})

	bot.Handle("/list", func(c telebot.Context) error {
		resp, err := http.Get("http://localhost:8080/api/v1/links")
		if err != nil {
			return err
		}

		var result linksResponse
		json.NewDecoder(resp.Body).Decode(&result)

		var message string
		for _, link := range result.Data {
			message += fmt.Sprintf("%s http://localhost:8080/%s (%d clicks)\n", link.OriginalURL, link.ShortCode, link.Clicks)
		}

		return c.Send(message)
	})

	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		url := c.Text()
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			return c.Send("invalid link")
		}

		body := strings.NewReader(`{"original_url": "` + url + `"}`)

		resp, err := http.Post("http://localhost:8080/api/v1/links", "application/json", body)
		if err != nil {
			return err
		}

		var result createResponse
		json.NewDecoder(resp.Body).Decode(&result)
		return c.Send("http://localhost:8080/" + result.Data.ShortCode)
	})

	bot.Start()
}
