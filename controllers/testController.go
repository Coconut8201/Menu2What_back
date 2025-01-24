package controllers

import (
	"Menu2What_back/interfaces"
	"Menu2What_back/utils/ApiResult"

	"github.com/gin-gonic/gin"
)

type TestController struct {
	interfaces.TestInterface
}

// Test handles the test endpoint
func (t *TestController) Test(c *gin.Context) {
	c.JSON(200, ApiResult.NewSuccessResult(200, gin.H{
		"message": "test success",
	}))
}

func (t *TestController) Aaa(c *gin.Context) {
	c.JSON(200, ApiResult.NewSuccessResult(200, gin.H{
		"message": "test success",
	}))
}
