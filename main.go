package main

import (
	"Menu2What_back/routers"
	"Menu2What_back/utils/Logger"
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

	logger, err := Logger.NewLogger(Logger.INFO)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("服務器啟動在 :6382")
	fmt.Println("http://localhost:6382/ping")
	if err := engine.Run(":6382"); err != nil {
		logger.Fatal("Server startup error: %v", err)
	}
}
