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
    baseRouter := router.Group("/api")
    apiRouter.Init("/api", testController, baseRouter)
    apiRouter.SetRoutes()

    return apiRouter
}

func (r *ApiRouter) Init(url string, controller interfaces.Controller, router *gin.RouterGroup) {
    fmt.Printf("初始化 ApiRouter，URL: %s\n", url)
    r.ctrl = controller
    r.BaseRoute.Init(url, controller, router)
}

func (r *ApiRouter) SetRoutes() {
    r.GetRoutes().GET("/test", func(c *gin.Context) {
        fmt.Println("收到請求：/api/test")
        r.ctrl.Test(c)
    })
}
