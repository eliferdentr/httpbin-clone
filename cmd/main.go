package main

import (
	"log"
	"os"

	router "httbinclone-eliferden.com/internal/httpbin"
)

func main() {
	// Router'ı oluştur
	r := router.NewRouter()

	// Port ayarla (ENV > default)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port :%s\n", port)

	// Server başlat
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
