package main

// @title HTTPBin Clone
// @version 1.0
// @description httpbin.org clone implemented with Go + Gin by Elif ERDEN.
// @BasePath /
// @schemes http

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
