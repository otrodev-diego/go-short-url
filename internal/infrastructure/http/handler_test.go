package http

import (
    "bytes"
    "dte-shortener/internal/application"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestShortenHandler(t *testing.T) {
    application.CleanupForTests()

    t.Run("Crear URL corta exitosamente", func(t *testing.T) {
        req := shortenRequest{
            DTEID:           "DTE001",
            ExpireInMinutes: 60,
        }
        body, _ := json.Marshal(req)

        request := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(body))
        response := httptest.NewRecorder()

        ShortenHandler(response, request)

        if response.Code != http.StatusOK {
            t.Errorf("Status code esperado %d, obtenido %d", http.StatusOK, response.Code)
        }

        var result shortenResponse
        err := json.Unmarshal(response.Body.Bytes(), &result)
        if err != nil {
            t.Fatalf("Error al decodificar respuesta: %v", err)
        }

        if result.ShortCode == "" {
            t.Error("La respuesta debería contener un código corto")
        }
    })

    t.Run("Request inválido", func(t *testing.T) {
        badJSON := []byte(`{malformed json}`)

        request := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(badJSON))
        response := httptest.NewRecorder()

        ShortenHandler(response, request)

        if response.Code != http.StatusBadRequest {
            t.Errorf("Status code esperado %d, obtenido %d", http.StatusBadRequest, response.Code)
        }
    })
}

func TestRedirectHandler(t *testing.T) {
    application.CleanupForTests()

    t.Run("Redirección exitosa", func(t *testing.T) {
        shortURL := application.CreateShortURL("DTE001", 60)

        request := httptest.NewRequest(http.MethodGet, "/s/"+shortURL.Code, nil)
        response := httptest.NewRecorder()

        RedirectHandler(response, request)

        if response.Code != http.StatusFound {
            t.Errorf("Status code esperado %d, obtenido %d", http.StatusFound, response.Code)
        }

        location := response.Header().Get("Location")
        expectedURL := "https://dte.example.com/DTE001"
        if location != expectedURL {
            t.Errorf("URL de redirección esperada %s, obtenida %s", expectedURL, location)
        }
    })

    t.Run("URL no encontrada", func(t *testing.T) {
        request := httptest.NewRequest(http.MethodGet, "/s/noexiste", nil)
        response := httptest.NewRecorder()

        RedirectHandler(response, request)

        if response.Code != http.StatusNotFound {
            t.Errorf("Status code esperado %d, obtenido %d", http.StatusNotFound, response.Code)
        }
    })

    t.Run("URL expirada", func(t *testing.T) {
        shortURL := application.CreateShortURL("DTE002", 0)

        request := httptest.NewRequest(http.MethodGet, "/s/"+shortURL.Code, nil)
        response := httptest.NewRecorder()

        RedirectHandler(response, request)

        if response.Code != http.StatusNotFound {
            t.Errorf("Status code esperado %d, obtenido %d", http.StatusNotFound, response.Code)
        }
    })
}