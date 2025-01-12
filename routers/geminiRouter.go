package routers

import (
	"Menu2What_back/controllers"
	"Menu2What_back/interfaces"

	"github.com/gin-gonic/gin"
)

type GeminiRouter struct {
	interfaces.BaseRoute
	ctrl interfaces.GeminiInterface
}

func NewGeminiRouter(router *gin.RouterGroup) *GeminiRouter {
	GeminiController := &controllers.GeminiController{}
	geminiRouter := &GeminiRouter{}
	baseRouter := router.Group("/gemini")
	geminiRouter.Init("/gemini", GeminiController, baseRouter)
	geminiRouter.SetRoutes()
	return geminiRouter
}

func (r *GeminiRouter) Init(url string, controller interfaces.GeminiInterface, router *gin.RouterGroup) {
	r.ctrl = controller
	r.BaseRoute.Init(url, controller, router)
}

func (r *GeminiRouter) SetRoutes() {
	r.GetRoutes().GET("/test", func(c *gin.Context) {
		r.ctrl.GeminiAPI(c)
	})
}
