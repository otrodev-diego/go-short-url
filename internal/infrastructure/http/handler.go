package http

import (
	"dte-shortener/internal/application"
	"dte-shortener/internal/infrastructure/security"
	"encoding/json"
	"net/http"
	"strings"
)

type shortenRequest struct {
	DTEID           string `json:"dte_id"`
	ExpireInMinutes int    `json:"expire_in_minutes"`
}

type shortenResponse struct {
    ShortCode string `json:"short_code"`
    Token     string `json:"token"`
}


func ShortenHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req shortenRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request format", http.StatusBadRequest)
        return
    }

    if req.DTEID == "" {
        http.Error(w, "DTE ID is required", http.StatusBadRequest)
        return
    }

    if req.ExpireInMinutes <= 0 {
        http.Error(w, "Expiration time must be greater than 0", http.StatusBadRequest)
        return
    }

    const maxExpirationMinutes = 24 * 60
    if req.ExpireInMinutes > maxExpirationMinutes {
        http.Error(w, "Expiration time cannot exceed 24 hours", http.StatusBadRequest)
        return
    }

    url := application.CreateShortURL(req.DTEID, req.ExpireInMinutes)

    token, err := security.GenerateToken(url.Code)
    if err != nil {
        http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
    }

    resp := shortenResponse{
        ShortCode: url.Code,
        Token:     token,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Path[len("/s/"):]

    token := r.Header.Get("Authorization")
    if token != "" {
        token = strings.TrimPrefix(token, "Bearer ")
        validCode, err := security.ValidateToken(token)
        if err != nil || validCode != code {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
    }

    url, ok := application.GetShortURL(code)
    if !ok {
        http.Error(w, "Link expired or not found", http.StatusNotFound)
        return
    }
    http.Redirect(w, r, url.TargetURL, http.StatusFound)
}

