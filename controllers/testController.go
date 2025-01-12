package controllers

import (
	"Menu2What_back/interfaces"
	"github.com/gin-gonic/gin"
)

type TestController struct {
	interfaces.TestInterface
}

// Test handles the test endpoint
func (t *TestController) Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test success",
	})
}

func (t *TestController) Aaa(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test success",
	})
}
