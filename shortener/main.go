package main

import (
	"log"
	"net/http"
	"shortener/handler"
	"shortener/storage"
)

func main() {
	memStorage := storage.NewMemoryStorage()

	h := handler.NewHandler(memStorage)

	// redisStorage := storage.NewRedisStorage("localhost:6379")
	// h := handler.NewHandler(redisStorage)

	//register route
	http.HandleFunc("POST /shorten", h.Shorten)
	http.HandleFunc("Get /{shortCode}", h.Redirect)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
