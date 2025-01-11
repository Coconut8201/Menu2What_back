package routers

import "github.com/gin-gonic/gin"

func SetupRouters(engine *gin.Engine) {
	NewTestRouter(engine.Group(""))
}
