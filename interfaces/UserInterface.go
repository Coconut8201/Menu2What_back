package interfaces

import (
	"github.com/gin-gonic/gin"
)

type UserInterface interface {
	Test(c *gin.Context)
	UserRegistration(c *gin.Context)
	UserLogin(c *gin.Context)
}
