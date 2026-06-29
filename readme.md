# URL Shortener with telegram bot

REST API for shortening URLs with a Telegram bot interface, built with Go and PostgreSQL.

## Stack

- **Go** — backend
- **PostgreSQL** — storage
- **Telegram Bot** — client interface
- **JWT** — authentication
- **Rate Limiting** — 3 req/sec per IP (token bucket algorithm)

## Getting Started

1. Clone the repository
2. Copy `.env.example` to `.env` and fill in your values
3. Run the server: `go run cmd/main.go`
4. Run the bot: `go run cmd/bot/main.go`

## API

### Public
| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/auth/telegram` | Get JWT token |
| `GET` | `/api/v1/links/{code}` | Get link info |
| `GET` | `/{code}` | Redirect to original URL |

### Protected (requires JWT)
| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/v1/links` | Create short link |
| `GET` | `/api/v1/links` | Get your links |
| `DELETE` | `/api/v1/links/{code}` | Delete link |

## Environment Variables

```env
DATABASE_URL=postgres://user:password@localhost:5432/urlshortener?sslmode=disable
SERVER_ADDR=:8080
TELEGRAM_TOKEN=your_telegram_bot_token
JWT_SECRET=your_secret_key
```

## Bot Commands

| Command | Description |
|---------|-------------|
| `/start` | Register and get started |
| `/list` | Show your links |
| `/delete {code}` | Delete a link |
| `/stats {code}` | Show link statistics |
| Send any URL | Returns shortened link |