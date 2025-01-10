package controllers

import (
    "github.com/gin-gonic/gin"
)

type TestController struct {
}

// Test handles the test endpoint
func (t *TestController) Test(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "test success",
    })
}
