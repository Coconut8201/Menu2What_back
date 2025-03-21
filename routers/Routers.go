package routers

import (
	"Menu2What_back/global"

	"github.com/gin-gonic/gin"
)

func SetupRouters(engine *gin.Engine) {
	NewTestRouter(engine.Group(""))
	NewGeminiRouter(engine.Group(""))
	NewUserRouter(engine.Group(""), global.DB)
	NewLineBotRouter(engine.Group(""))
}
