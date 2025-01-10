package main

import (
    "Menu2What_back/routers"
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
    engine := gin.Default()

    routers.NewApiRouter(engine.Group(""))

    engine.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    fmt.Println("服務器啟動在 :6382")
    engine.Run(":6382")
}
