package router

import (
	"API/internal/handler"
	"API/internal/middleware"
	"API/internal/repository"
	"API/internal/service"
	"database/sql"
	"net/http"
)

func SetupRoutes(db *sql.DB, secret string) http.Handler {
	authHandler := handler.NewAuthHandler(secret)
	rateLimiter := middleware.NewRateLimiter(3, 5)
	authMiddleware := middleware.NewAuthMiddleware(secret)
	linkRepo := repository.NewPostgresLinkRepo(db)
	linkService := service.NewLinkService(linkRepo)
	linkHandler := handler.NewLinkHandler(linkService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /auth/telegram", authHandler.Auth)
	mux.Handle("POST /api/v1/links", authMiddleware.Require(http.HandlerFunc(linkHandler.Create)))
	mux.Handle("GET /api/v1/links", authMiddleware.Require(http.HandlerFunc(linkHandler.GetAll)))
	mux.HandleFunc("GET /api/v1/links/{code}", linkHandler.GetByCode)
	mux.Handle("DELETE /api/v1/links/{code}", authMiddleware.Require(http.HandlerFunc(linkHandler.Delete)))
	mux.HandleFunc("GET /{code}", linkHandler.Redirect)

	return rateLimiter.Limit(mux)
}
