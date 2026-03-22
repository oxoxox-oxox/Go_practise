package handler

import (
	"encoding/json"
	"net/http"
	"shortener/storage"
	"shortener/utils"
)

type Handler struct {
	storage storage.Storage
}

func NewHandler(s storage.Storage) *Handler {
	return &Handler{storage: s}
}

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	ShortCode string 
	ShortURL string
}


//POST /shorten

func(h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}
	//generate shortcode
	shortCode := utils.GenerateShortCode()
	if err := h.storage.Save(shortCode, req.URL); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	shortURL := scheme + "://" + host + "/" + shortCode

	resp := shortenResponse{
		ShortCode: shortCode,
		ShortURL: shortURL,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

//Get method
func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:]
	longURL, ok := h.storage.Load(shortCode)
	if !ok {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, longURL, http.StatusFound)
}


