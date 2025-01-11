package controllers

import (
	"github.com/gin-gonic/gin"
	"Menu2What_back/interfaces"
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
