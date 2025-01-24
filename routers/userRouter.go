package routers

import (
	"Menu2What_back/controllers"
	"Menu2What_back/interfaces"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRouter struct {
	interfaces.BaseRoute
	ctrl interfaces.UserInterface
}

func NewUserRouter(router *gin.RouterGroup, db *gorm.DB) *UserRouter {
	UserController := &controllers.UserController{
		DB: db,
	}
	userRouter := &UserRouter{}
	baseRouter := router.Group("/user")
	userRouter.Init("/user", UserController, baseRouter)
	userRouter.SetRoutes()
	return userRouter
}

func (r *UserRouter) Init(url string, controller interfaces.UserInterface, router *gin.RouterGroup) {
	r.ctrl = controller
	r.BaseRoute.Init(url, controller, router)
}

func (r *UserRouter) SetRoutes() {
	r.GetRoutes().GET("/test", func(c *gin.Context) {
		r.ctrl.Test(c)
	})

	// 添加註冊路由
	r.GetRoutes().POST("/register", func(c *gin.Context) {
		r.ctrl.UserRegistration(c)
	})

	r.GetRoutes().POST("/token", func(c *gin.Context) {
		r.ctrl.UserToken(c)
	})

	r.GetRoutes().POST("/logout", func(c *gin.Context) {
		r.ctrl.UserLogout(c)
	})
}
