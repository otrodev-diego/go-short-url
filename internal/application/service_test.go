package application

import (
    "testing"
    "time"
    "dte-shortener/internal/domain"
)

func TestCreateShortURL(t *testing.T) {
    shortUrls = make(map[string]*domain.ShortURL)

    t.Run("Crear URL corta", func(t *testing.T) {
        dteID := "DTE001"
        expireMinutes := 60

        shortURL := CreateShortURL(dteID, expireMinutes)

        if shortURL == nil {
            t.Fatal("Se esperaba una URL corta, se obtuvo nil")
        }
        if shortURL.DTEID != dteID {
            t.Errorf("DTEID esperado %s, obtenido %s", dteID, shortURL.DTEID)
        }
        if shortURL.TargetURL != "https://dte.example.com/"+dteID {
            t.Errorf("TargetURL esperada %s, obtenida %s", "https://dte.example.com/"+dteID, shortURL.TargetURL)
        }
        if shortURL.Code == "" {
            t.Error("El código generado no debería estar vacío")
        }
        if len(shortURL.Code) != 6 {
            t.Errorf("La longitud del código debería ser 6, obtenido %d", len(shortURL.Code))
        }
    })
}

func TestGetShortURL(t *testing.T) {
    shortUrls = make(map[string]*domain.ShortURL)

    t.Run("Obtener URL existente y no expirada", func(t *testing.T) {
        dteID := "DTE002"
        expireMinutes := 60
        created := CreateShortURL(dteID, expireMinutes)

        found, exists := GetShortURL(created.Code)
        if !exists {
            t.Error("La URL debería existir")
        }
        if found == nil {
            t.Fatal("Se esperaba encontrar la URL, se obtuvo nil")
        }
        if found.Code != created.Code {
            t.Errorf("Código esperado %s, obtenido %s", created.Code, found.Code)
        }
        if found.AccessCount != 1 {
            t.Errorf("AccessCount esperado 1, obtenido %d", found.AccessCount)
        }
    })

    t.Run("Obtener URL no existente", func(t *testing.T) {
        _, exists := GetShortURL("codigo-no-existente")
        if exists {
            t.Error("No debería existir la URL")
        }
    })

    t.Run("Obtener URL expirada", func(t *testing.T) {
        dteID := "DTE003"
        expireMinutes := 0
        created := CreateShortURL(dteID, expireMinutes)

        time.Sleep(time.Millisecond)

        _, exists := GetShortURL(created.Code)
        if exists {
            t.Error("La URL expirada no debería ser accesible")
        }
    })
}

func TestGenerateCode(t *testing.T) {
    code1 := generateCode(6)
    code2 := generateCode(6)

    if len(code1) != 6 {
        t.Errorf("La longitud del código debería ser 6, obtenido %d", len(code1))
    }
    if code1 == code2 {
        t.Error("Los códigos generados no deberían ser iguales")
    }

    validChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    for _, char := range code1 {
        if !contains(validChars, char) {
            t.Errorf("Caracter inválido encontrado: %c", char)
        }
    }
}

func contains(s string, c rune) bool {
    for _, char := range s {
        if char == c {
            return true
        }
    }
    return false
}