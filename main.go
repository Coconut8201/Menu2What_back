package main

import (
	"Menu2What_back/routers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	routers.SetupRouters(engine)

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server is running",
		})
	})

	fmt.Println("服務器啟動在 :6382")
	fmt.Println("http://localhost:6382/ping")
	if err := engine.Run(":6382"); err != nil {
		log.Fatal("Server startup error:", err)
	}
}
