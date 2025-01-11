package routers

import (
	"Menu2What_back/controllers"
	"Menu2What_back/interfaces"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TestRouter struct {
	interfaces.BaseRoute
	ctrl interfaces.TestInterface
}

func NewTestRouter(router *gin.RouterGroup) *TestRouter {
	testController := &controllers.TestController{}
	testRouter := &TestRouter{}
	baseRouter := router.Group("/api")
	testRouter.Init("/apis", testController, baseRouter)
	testRouter.SetRoutes()
	return testRouter
}

func (r *TestRouter) Init(url string, controller interfaces.TestInterface, router *gin.RouterGroup) {
	r.ctrl = controller
	r.BaseRoute.Init(url, controller, router)
}

func (r *TestRouter) SetRoutes() {
	r.GetRoutes().GET("/test", func(c *gin.Context) {
		fmt.Println("收到請求：/api/test")
		r.ctrl.Test(c)
	})
}
