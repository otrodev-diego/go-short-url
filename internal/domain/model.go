package domain

import "time"

type ShortURL struct {
	Code        string
	DTEID       string
	TargetURL   string
	ExpiresAt   time.Time
	AccessCount int
}
