package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello, World!")
	})

	err := router.Run(":4321")
	if err != nil {
		log.Fatal("Error running server, ", err)
	}
}
