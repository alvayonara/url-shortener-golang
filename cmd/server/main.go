package main

import (
	"fmt"
	"log"
	"net/http"
	httpHandler "url-shortener-golang/internal/http"
	"url-shortener-golang/internal/shortener"
	"url-shortener-golang/internal/store"
)

func main() {
	store, err := store.NewMySQLStore(
		"root:root@tcp(localhost:3306)/url_shortener",
	)
	if err != nil {
		log.Fatal(err)
	}
	//store := store.NewMemoryStore() --> in-memory approach
	service := shortener.NewService(store)
	handler := httpHandler.NewHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/shorten", handler.Shorten)
	mux.HandleFunc("/", handler.Redirect)

	fmt.Println("Server running properly.")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
