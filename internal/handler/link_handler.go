package handler

import (
	"API/internal/models"
	"API/internal/service"
	"encoding/json"
	"net/http"
)

type LinkHandler struct {
	service *service.LinkService
}

func NewLinkHandler(service *service.LinkService) *LinkHandler {
	return &LinkHandler{service: service}
}

type CreateLinkRequest struct {
	OriginalURL string `json:"original_url"`
}

type LinkResponse struct {
	Data *models.Link `json:"data"`
}

type LinksResponse struct {
	Data []*models.Link `json:"data"`
}

func (h *LinkHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateLinkRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	link, err := h.service.CreateLink(r.Context(), req.OriginalURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(LinkResponse{Data: link})
}

func (h *LinkHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	links, err := h.service.GetAllLinks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LinksResponse{Data: links})
}

func (h *LinkHandler) GetByCode(w http.ResponseWriter, r *http.Request) {

	code := r.PathValue("code")

	link, err := h.service.GetByCode(r.Context(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LinkResponse{Data: link})
}

func (h *LinkHandler) Delete(w http.ResponseWriter, r *http.Request) {

	code := r.PathValue("code")

	err := h.service.DeleteLink(r.Context(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *LinkHandler) Redirect(w http.ResponseWriter, r *http.Request) {

	code := r.PathValue("code")

	link, err := h.service.GetByCode(r.Context(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_ = h.service.IncrementClicks(r.Context(), code)

	http.Redirect(w, r, link.OriginalURL, http.StatusMovedPermanently)
}
