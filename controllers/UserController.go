package controllers

import (
	"Menu2What_back/interfaces"
	"Menu2What_back/services/user"
	"Menu2What_back/utils/ApiResult"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	interfaces.UserInterface
	DB *gorm.DB
}

// UserRegistrationRequest 使用者註冊請求
type UserRegistrationRequest struct {
	UserName     string `json:"userName" binding:"required" example:"admin"`
	UserPassword string `json:"userPassword" binding:"required" example:"password"`
}

// UserRegistrationErrorResponse 錯誤回應結構
type UserRegistrationErrorResponse struct {
	Error  string `json:"error" example:"處理請求時發生錯誤"`
	Detail string `json:"detail,omitempty" example:"無效的輸入格式"`
}

// UserLoginRequest 使用者登入請求
type UserLoginRequest struct {
	UserName     string `json:"userName" binding:"required" example:"admin"`
	UserPassword string `json:"userPassword" binding:"required" example:"password"`
}

// UserLoginErrorResponse 登入錯誤回應
type UserLoginErrorResponse struct {
	Error  string `json:"error" example:"登入失敗"`
	Detail string `json:"detail,omitempty" example:"無效的憑證"`
}

// Test godoc 測試用戶路徑
// @Summary 測試用戶控制器
// @Description 測試用戶控制器是否正常運作
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "返回成功訊息"
// @Router /api/user/test [get]
func (u *UserController) Test(c *gin.Context) {
	c.JSON(200, ApiResult.NewSuccessResult(200, gin.H{
		"message": "user success",
	}))
}

// UserRegistration godoc 使用者註冊
// @Summary 用戶註冊
// @Description 創建新用戶帳號
// @Tags User
// @Accept json
// @Produce json
// @Param request body UserRegistrationRequest true "用戶註冊資訊"
// @Success 200 {object} map[string]interface{} "註冊成功"
// @Failure 400 {object} UserRegistrationErrorResponse "註冊失敗"
// @Router /api/user/register [post]
func (u *UserController) UserRegistration(c *gin.Context) {
	var req UserRegistrationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, ApiResult.NewFailResult(400, "使用者資訊錯誤："+err.Error()))
		return
	}

	if err := user.CreateUser(u.DB, req.UserName, req.UserPassword); err != nil {
		c.JSON(400, ApiResult.NewFailResult(400, "註冊失敗："+err.Error()))
		return
	}

	c.JSON(200, ApiResult.NewSuccessResult(200, gin.H{
		"message": "註冊成功",
	}))
}

// UserToken godoc 使用者登入
// @Summary 用戶登入
// @Description 使用帳號密碼進行登入
// @Tags User
// @Accept json
// @Produce json
// @Param request body UserLoginRequest true "用戶登入資訊"
// @Success 200 {object} map[string]interface{} "登入成功"
// @Failure 400 {object} UserLoginErrorResponse "登入失敗"
// @Router /api/user/login [post]
func (u *UserController) UserToken(c *gin.Context) {
	var req UserLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, ApiResult.NewFailResult(400, "登入資訊錯誤："+err.Error()))
		return
	}

	authenticatedUser, err := user.Authenticate(u.DB, req.UserName, req.UserPassword)
	if err != nil {
		c.JSON(400, ApiResult.NewFailResult(400, "登入失敗：無效的帳號或密碼"))
		return
	}

	token, err := authenticatedUser.GenerateJWT()
	if err != nil {
		c.JSON(500, ApiResult.NewFailResult(500, "登入失敗：無法產生授權token"))
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "", false, true)

	c.JSON(200, ApiResult.NewSuccessResult(200, gin.H{
		"message": "登入成功",
		"token":   token,
	}))
}

func (u *UserController) UserLogout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, ApiResult.NewSuccessResult(200, gin.H{
		"message": "成功登出",
	}))
}
