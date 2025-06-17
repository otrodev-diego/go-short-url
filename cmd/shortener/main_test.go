package main

import (
    "net/http"
    "testing"
    "time"
)

func TestServer(t *testing.T) {
    go main()

    time.Sleep(100 * time.Millisecond)

    t.Run("Servidor responde", func(t *testing.T) {
        resp, err := http.Get("http://localhost:8081/health")
        if err != nil {
            t.Fatalf("Error al hacer request: %v", err)
        }
        defer resp.Body.Close()

        if resp.StatusCode != http.StatusOK {
            t.Errorf("código de estado esperado %d, obtenido %d", http.StatusOK, resp.StatusCode)
        }
    })
}