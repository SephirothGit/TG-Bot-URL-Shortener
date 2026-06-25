# URL Shortener with telegram bot

REST API for shortening URLs with a Telegram bot interface, built with Go and PostgreSQL.

## Stack

- **Go** — backend
- **PostgreSQL** — storage
- **Telegram Bot** — client interface

## Getting Started

1. Clone the repository
2. Copy `.env.example` to `.env` and fill in your values
3. Run the server: `go run cmd/main.go`
4. Run the bot: `go run cmd/bot/main.go`

## API

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/v1/links` | Create short link |
| `GET` | `/api/v1/links` | Get all links |
| `GET` | `/api/v1/links/{code}` | Get link by code |
| `DELETE` | `/api/v1/links/{code}` | Delete link |
| `GET` | `/{code}` | Redirect to original URL |

## Environment Variables

```env
DATABASE_URL=postgres://user:password@localhost:5432/urlshortener?sslmode=disable
SERVER_ADDR=:8080
TELEGRAM_TOKEN=your_telegram_bot_token
```

## Bot Commands

| Command | Description |
|---------|-------------|
| `/start` | Welcome message |
| Send any URL | Returns shortened link |