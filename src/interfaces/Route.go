package interfaces

import (
	"github.com/gin-gonic/gin"
)

type Route interface {
	setRoutes()
	GetRoutes() *gin.RouterGroup
	GetUrl() string
}

type BaseRoute struct {
	url    string
	ctrl   Controller
	router *gin.RouterGroup
}

func (r *BaseRoute) GetRouter() *gin.RouterGroup {
	return r.router
}

func (r *BaseRoute) GetURL() string {
	return r.url
}
