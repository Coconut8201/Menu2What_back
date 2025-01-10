package routers

import (
	"Menu2What_back/src/controllers"
	"Menu2What_back/src/interfaces"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	interfaces.BaseRoute
	ctrl interfaces.Controller
}

func NewApiRouter(router *gin.RouterGroup) *ApiRouter {
	fmt.Println("=== 創建新的 ApiRouter ===")
	testController := &controllers.TestController{}
	apiRouter := &ApiRouter{}

	// 創建基礎路由組
	baseRouter := router.Group("/api")

	// 初始化
	apiRouter.Init("/api", testController, baseRouter)

	// 確保設置路由
	apiRouter.SetRoutes()

	return apiRouter
}

func (r *ApiRouter) Init(url string, controller interfaces.Controller, router *gin.RouterGroup) {
	fmt.Printf("初始化 ApiRouter，URL: %s\n", url)
	r.ctrl = controller
	r.BaseRoute.Init(url, controller, router)
}

func (r *ApiRouter) SetRoutes() {
	fmt.Println("=== 開始設置路由 ===")
	fmt.Printf("URL: %s\n", r.GetURL())

	if r.GetRoutes() == nil {
		fmt.Println("警告：router 是 nil")
		return
	}

	if r.ctrl == nil {
		fmt.Println("警告：controller 是 nil")
		return
	}

	r.GetRoutes().GET("/test", func(c *gin.Context) {
		fmt.Println("收到請求：/api/test")
		r.ctrl.Test(c)
	})

	fmt.Println("=== 路由設置完成 ===")
}
