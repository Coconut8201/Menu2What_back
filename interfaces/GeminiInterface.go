package interfaces

import (
	"github.com/gin-gonic/gin"
)

type GeminiInterface interface {
	Test(c *gin.Context)
	GeminiApiTextOnly(c *gin.Context)
	GeminiApiImageAndText(c *gin.Context)
}
