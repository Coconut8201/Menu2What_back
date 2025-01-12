package main

import (
	_ "Menu2What_back/docs"
	"Menu2What_back/routers"
	"Menu2What_back/utils/Logger"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Menu2What API
// @version         1.0
// @description     這是一個 Menu2What 的 API 服務
// @host      localhost:6382
// @BasePath  /
func main() {
	engine := gin.Default()

	// Swagger 文檔路由要在其他路由之前
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:6382/swagger/doc.json")))

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
