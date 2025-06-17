package main

import (
    "log"
    "net/http"
    httphandler "dte-shortener/internal/infrastructure/http"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    })

    mux.HandleFunc("/shorten", httphandler.ShortenHandler)
    mux.HandleFunc("/s/", httphandler.RedirectHandler)

    server := &http.Server{
        Addr:    ":8081",
        Handler: mux,
    }

    log.Println("Server running on http://localhost:8081")
    log.Fatal(server.ListenAndServe())
}