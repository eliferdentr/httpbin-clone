package main

// @title HTTPBin Clone
// @version 1.0
// @description router.org clone implemented with Go + Gin by Elif ERDEN.
// @BasePath /
// @schemes http

import (
	"log"
	"os"

	_ "httbinclone-eliferden.com/docs"
	router "httbinclone-eliferden.com/internal/router"
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
