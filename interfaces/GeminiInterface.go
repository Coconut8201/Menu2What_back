package interfaces

import "github.com/gin-gonic/gin"

type GeminiInterface interface {
	Test(c *gin.Context)
	GeminiAPI(c *gin.Context)
}
