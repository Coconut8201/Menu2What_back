package routers

import (
	"Menu2What_back/controllers"
	"Menu2What_back/interfaces"

	"github.com/gin-gonic/gin"
)

type LineBotRouter struct {
	interfaces.BaseRoute
	ctrl interfaces.LineBotInterface
}

func NewLineBotRouter(router *gin.RouterGroup) *LineBotRouter {
	lineBotController := &controllers.LineBotController{}
	lineBotRouter := &LineBotRouter{}
	baseRouter := router.Group("/linebot")
	lineBotRouter.Init("/linebot", lineBotController, baseRouter)
	lineBotRouter.SetRoutes()
	return lineBotRouter
}

func (r *LineBotRouter) Init(url string, controller interfaces.LineBotInterface, router *gin.RouterGroup) {
	r.ctrl = controller
	r.BaseRoute.Init(url, controller, router)
}

func (r *LineBotRouter) SetRoutes() {
	r.GetRoutes().GET("/test", func(c *gin.Context) {
		r.ctrl.Test(c)
	})
}