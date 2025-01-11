package interfaces

import "github.com/gin-gonic/gin"

type TestInterface interface {
	Test(c *gin.Context)
	Aaa(c *gin.Context)
}
