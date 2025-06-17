package domain

import (
    "testing"
    "time"
)

func TestShortURL(t *testing.T) {
    t.Run("Creación de ShortURL", func(t *testing.T) {
        now := time.Now()
        shortURL := ShortURL{
            Code:        "abc123",
            DTEID:       "DTE001",
            TargetURL:   "https://dte.example.com/document/123",
            ExpiresAt:   now.Add(24 * time.Hour),
            AccessCount: 0,
        }

        if shortURL.Code != "abc123" {
            t.Errorf("Code esperado abc123, obtenido %s", shortURL.Code)
        }
        if shortURL.DTEID != "DTE001" {
            t.Errorf("DTEID esperado DTE001, obtenido %s", shortURL.DTEID)
        }
        if shortURL.TargetURL != "https://dte.example.com/document/123" {
            t.Errorf("TargetURL esperada https://dte.example.com/document/123, obtenido %s", shortURL.TargetURL)
        }
        if shortURL.AccessCount != 0 {
            t.Errorf("AccessCount inicial esperado 0, obtenido %d", shortURL.AccessCount)
        }
    })

    t.Run("Verificación de expiración", func(t *testing.T) {
        notExpired := ShortURL{
            Code:      "abc123",
            ExpiresAt: time.Now().Add(1 * time.Hour),
        }
        if notExpired.ExpiresAt.Before(time.Now()) {
            t.Error("La URL no debería estar expirada")
        }

        expired := ShortURL{
            Code:      "xyz789",
            ExpiresAt: time.Now().Add(-1 * time.Hour),
        }
        if !expired.ExpiresAt.Before(time.Now()) {
            t.Error("La URL debería estar expirada")
        }
    })

    t.Run("Control de accesos", func(t *testing.T) {
        shortURL := ShortURL{
            Code:        "test123",
            AccessCount: 0,
        }

        initialCount := shortURL.AccessCount
        shortURL.AccessCount++

        if shortURL.AccessCount != initialCount+1 {
            t.Errorf("AccessCount esperado %d, obtenido %d", initialCount+1, shortURL.AccessCount)
        }
    })
}