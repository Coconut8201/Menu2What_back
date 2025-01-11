package interfaces

import "github.com/gin-gonic/gin"

type TestController interface {
	Controller
	Test(c *gin.Context)
	Aaa(c *gin.Context)
}
