package main

import (
	"encoding/json"
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

func main() {
	godotenv.Load()

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

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
