package main

import (
	"github.com/gin-gonic/gin"
	db "httbinclone-eliferden.com/db"
	router "httbinclone-eliferden.com/routes"
)

func main() {
	server := gin.Default()
	router.RegisterRoutes(server)
	db.InitDB()
}