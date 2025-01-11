package interfaces

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Test(c *gin.Context)
}
