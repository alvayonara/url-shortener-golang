package http

import (
	"encoding/json"
	"net/http"
	"url-shortener-golang/internal/shortener"
)

type Handler struct {
	service *shortener.Service
}

type shortenRequest struct {
	URL string `json:"url""`
}

type shortenResponse struct {
	Code string `json:"code"`
	URL  string `json:"url"`
}

func NewHandler(service *shortener.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}
	if req.URL == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}
	link, err := h.service.Create(req.URL)
	if err != nil {
		http.Error(w, "failed to create short url", http.StatusInternalServerError)
		return
	}
	resp := shortenResponse{
		Code: link.Code,
		URL:  link.URL,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
