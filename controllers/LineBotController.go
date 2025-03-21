package controllers

import (
	"Menu2What_back/interfaces"
	"Menu2What_back/utils/ApiResult"

	"github.com/gin-gonic/gin"
)

type LineBotController struct {
	interfaces.LineBotInterface
}

func(g * LineBotController) Test(c *gin.Context){
	c.JSON(200, ApiResult.NewSuccessResult(200, gin.H{
		"message": "linebot success",
	}))
}

func (g * LineBotController) LineBotEcho(c *gin.Context){
	c.JSON(200, ApiResult.NewSuccessResult(200, gin.H{
		"message": "linebot success",
	}))
}


