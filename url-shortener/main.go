package main

import (
	"fmt"
	"log"
	"url-shortener/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db, err := config.DBConnect()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
