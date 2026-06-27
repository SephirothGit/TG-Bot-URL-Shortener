package models

import "time"

type Link struct {
	ID          int       `json:"id"`
	TelegramID  int64     `json:"telegram_id"`
	OriginalURL string    `json:"original_url"`
	ShortCode   string    `json:"short_code"`
	Clicks      int       `json:"clicks"`
	CreatedAt   time.Time `json:"created_at"`
}
