package interfaces

import (
    "github.com/gin-gonic/gin"
)

type Route interface {
    SetRoutes()
    GetRoutes() *gin.RouterGroup
    GetURL() string
}

type BaseRoute struct {
    url    string
    ctrl   Controller
    router *gin.RouterGroup
}

func (b *BaseRoute) GetURL() string {
    return b.url
}

func (b *BaseRoute) SetRoutes() {}

func (b *BaseRoute) GetRoutes() *gin.RouterGroup {
    return b.router
}

func (b *BaseRoute) Init(url string, ctrl Controller, router *gin.RouterGroup) {
    b.url = url
    b.ctrl = ctrl
    b.router = router
    b.SetRoutes()
}

func NewBaseRoute(url string, ctrl Controller, router *gin.RouterGroup) *BaseRoute {
    return &BaseRoute{
        url:    url,
        ctrl:   ctrl,
        router: router,
    }
}
