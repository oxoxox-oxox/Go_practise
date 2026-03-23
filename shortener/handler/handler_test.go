package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"shortener/storage"
	"testing"
)

func TestSshortenAndRedirect(t *testing.T) {
	memStorage := storage.NewMemoryStorage()
	h := NewHandler(memStorage)

	//test create shortener
	reqBody := shortenRequest{URL: "https://example.com"}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/shorten", bytes.NewReader(body))
	w := httptest.NewRecorder()
	h.Shorten(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status 201, git %d", resp.StatusCode)
	}

	var respData shortenResponse
	json.NewDecoder(resp.Body).Decode(&respData)
	if respData.ShortCode == "" {
		t.Error("expected short code")
	}

	req2 := httptest.NewRequest("GET", "/"+respData.ShortCode, nil)
	w2 := httptest.NewRecorder()
	h.Redirect(w2, req2)

	if w2.Code != http.StatusFound {
		t.Errorf("expected redirect 302, got %d", w2.Code)
	}
	location := w2.Header().Get("Location")
	if location != "https://example.com" {
		t.Errorf("expected location https//example.com, got %s", location)
	}
}
