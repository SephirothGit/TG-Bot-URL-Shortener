package router

import (
	"API/internal/handler"
	"API/internal/repository"
	"API/internal/service"
	"database/sql"
	"net/http"
)

func SetupRoutes(db *sql.DB) *http.ServeMux {
	linkRepo := repository.NewPostgresLinkRepo(db)
	linkService := service.NewLinkService(linkRepo)
	linkHandler := handler.NewLinkHandler(linkService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/links", linkHandler.Create)
	mux.HandleFunc("GET /api/v1/links", linkHandler.GetAll)
	mux.HandleFunc("GET /api/v1/links/{code}", linkHandler.GetByCode)
	mux.HandleFunc("DELETE /api/v1/links/{code}", linkHandler.Delete)
	mux.HandleFunc("GET /{code}", linkHandler.Redirect)

	return mux
}
