package domain

import "time"

// Estructura de una URL corta
type ShortURL struct {
	Code        string
	DTEID       string
	TargetURL   string
	ExpiresAt   time.Time
	AccessCount int
}
