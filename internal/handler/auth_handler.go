package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthRequest struct {
	TelegramID int64 `json:"telegram_id"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type AuthHandler struct {
	secret string
}

func NewAuthHandler(secret string) *AuthHandler {
	return &AuthHandler{secret: secret}
}

func (h *AuthHandler) Auth(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	claims := jwt.MapClaims{
		"telegram_id": req.TelegramID,
		"exp":         time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(h.secret))
	if err != nil {
		http.Error(w, "failed to create token", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthResponse{Token: tokenString})
}
