package application

import (
	"dte-shortener/internal/domain"
	"math/rand"
	"time"
)

var shortUrls = make(map[string]*domain.ShortURL)

func generateCode(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func CreateShortURL(dteID string, expireMinutes int) *domain.ShortURL {
	code := generateCode(6)
	url := &domain.ShortURL{
		Code:      code,
		DTEID:     dteID,
		TargetURL: "https://dte.example.com/" + dteID,
		ExpiresAt: time.Now().Add(time.Duration(expireMinutes) * time.Minute),
	}
	shortUrls[code] = url
	return url
}

func GetShortURL(code string) (*domain.ShortURL, bool) {
	url, exists := shortUrls[code]
	if !exists || time.Now().After(url.ExpiresAt) {
		return nil, false
	}
	url.AccessCount++
	return url, true
}

func CleanupForTests() {
    shortUrls = make(map[string]*domain.ShortURL)
}

