package main

import (
	"github.com/gin-gonic/gin"
	db "httbinclone-eliferden.com/db"
)

func main() {
	server := gin.Default()
	
	db.InitDB()
}