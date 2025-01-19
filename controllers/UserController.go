package controllers

import (
	"Menu2What_back/interfaces"
	"Menu2What_back/services/user"

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

// Test godoc 測試用戶路徑
// @Summary 測試用戶控制器
// @Description 測試用戶控制器是否正常運作
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "返回成功訊息"
// @Router /api/user/test [get]
func (u *UserController) Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user success",
	})
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
		c.JSON(400, UserRegistrationErrorResponse{
			Error:  "使用者資訊錯誤",
			Detail: err.Error(),
		})
		return
	}

	// 呼叫創建使用者服務
	if err := user.CreateUser(u.DB, req.UserName, req.UserPassword); err != nil {
		c.JSON(400, UserRegistrationErrorResponse{
			Error:  "註冊失敗",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "註冊成功",
	})
}
