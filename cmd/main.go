package main

import (
	"github.com/gin-gonic/gin"
	router "httbinclone-eliferden.com/routes"
)

func main() {
	server := gin.Default()
	router.RegisterRoutes(server)
	server.Run()
}