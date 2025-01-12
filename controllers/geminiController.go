package controllers

import (
	"Menu2What_back/interfaces"
	"Menu2What_back/utils/tools"

	"github.com/gin-gonic/gin"
)

type GeminiController struct {
	interfaces.GeminiInterface
}

func (g *GeminiController) Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "gemini success",
	})
}

// GeminiAPIRequest 請求結構
type GeminiAPIRequest struct {
	Message string `json:"message" binding:"required" example:"你好"`
}

// GeminiAPIResponse 成功回應結構
type GeminiAPIResponse struct {
	Message string `json:"message" example:"你好，我是 Gemini AI"`
}

// GeminiAPIErrorResponse 錯誤回應結構
type GeminiAPIErrorResponse struct {
	Error  string `json:"error" example:"處理請求時發生錯誤"`
	Detail string `json:"detail,omitempty" example:"無效的輸入格式"`
}

// GeminiAPI godoc
// @Summary      Gemini AI API
// @Description  傳送訊息到 Gemini AI 並獲取回應
// @Tags         gemini
// @Accept       json
// @Produce      json
// @Param        request body GeminiAPIRequest true "請求內容"
// @Success      200  {object}  GeminiAPIResponse
// @Failure      400  {object}  GeminiAPIErrorResponse
// @Failure      500  {object}  GeminiAPIErrorResponse
// @Router       /api [post]
func (g *GeminiController) GeminiAPI(c *gin.Context) {
	var req GeminiAPIRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, GeminiAPIErrorResponse{
			Error:  "無法解析請求數據",
			Detail: err.Error(),
		})
		return
	}

	resultChan := make(chan struct {
		response string
		err      error
	})

	go func() {
		resp, err := tools.CallGeminiApi(req.Message)
		resultChan <- struct {
			response string
			err      error
		}{resp, err}
	}()

	result := <-resultChan
	if result.err != nil {
		c.JSON(500, GeminiAPIErrorResponse{
			Error: result.err.Error(),
		})
		return
	}

	c.JSON(200, GeminiAPIResponse{
		Message: result.response,
	})
}
