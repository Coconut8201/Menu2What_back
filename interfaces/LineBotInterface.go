package interfaces

import (
	"github.com/gin-gonic/gin"
)

type LineBotInterface interface {
	Test(c *gin.Context)
	LineBotEcho(c *gin.Context)
}

func LineBotEcho(c *gin.Context){

}