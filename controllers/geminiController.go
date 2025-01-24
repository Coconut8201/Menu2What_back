package controllers

import (
	"Menu2What_back/interfaces"
	"Menu2What_back/utils/ApiResult"
	"Menu2What_back/utils/tools"

	"io"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type GeminiController struct {
	interfaces.GeminiInterface
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

// GeminiAPIRequestImageAndText 同時上傳圖片文字請求結構
type GeminiAPIRequestImageAndText struct {
	Message string                `form:"message" binding:"required" example:"請描述這張圖片"`
	Image   *multipart.FileHeader `form:"image" binding:"required" example:"image.jpg"`
}

func (g *GeminiController) Test(c *gin.Context) {
	c.JSON(200, ApiResult.NewSuccessResult(200, gin.H{
		"message": "gemini success",
	}))
}

// GeminiApiTextOnly godoc
// @Summary      Gemini AI API
// @Description  傳送訊息到 Gemini AI 並獲取回應
// @Tags         gemini
// @Accept       json
// @Produce      json
// @Param        request body GeminiAPIRequest true "請求內容"
// @Success      200  {object}  GeminiAPIResponse
// @Failure      400  {object}  GeminiAPIErrorResponse
// @Failure      500  {object}  GeminiAPIErrorResponse
// @Router       /gemini/api/text_only [post]
func (g *GeminiController) GeminiApiTextOnly(c *gin.Context) {
	var req GeminiAPIRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, ApiResult.NewFailResult(400, "無法解析請求數據"))
		return
	}

	resultChan := make(chan struct {
		response string
		err      error
	})

	go func() {
		resp, err := tools.CallGeminiApiTextOnly(req.Message)
		resultChan <- struct {
			response string
			err      error
		}{resp, err}
	}()

	result := <-resultChan
	if result.err != nil {
		c.JSON(500, ApiResult.NewFailResult(500, result.err.Error()))
		return
	}

	c.JSON(200, ApiResult.NewSuccessResult(200, gin.H{
		"message": result.response,
	}))
}

// GeminiApiImageAndText godoc
// @Summary      Gemini AI 圖文分析 API
// @Description  上傳圖片和文字到 Gemini AI 並獲取分析回應
// @Tags         gemini
// @Accept       multipart/form-data
// @Produce      json
// @Param        message formData string true "文字訊息"
// @Param        image formData file true "圖片檔案"
// @Success      200  {object}  GeminiAPIResponse
// @Failure      400  {object}  GeminiAPIErrorResponse
// @Failure      500  {object}  GeminiAPIErrorResponse
// @Router       /gemini/api/image_and_text [post]
func (g *GeminiController) GeminiApiImageAndText(c *gin.Context) {
	var req GeminiAPIRequestImageAndText
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, ApiResult.NewFailResult(400, "無法解析請求數據"))
		return
	}

	if req.Image == nil {
		c.JSON(400, ApiResult.NewFailResult(400, "未提供圖片"))
		return
	}

	resultChan := make(chan struct {
		response string
		err      error
	})

	file, err := req.Image.Open()
	if err != nil {
		c.JSON(400, ApiResult.NewFailResult(400, "無法開啟圖片檔案"))
		return
	}
	defer file.Close()

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		c.JSON(400, ApiResult.NewFailResult(400, "無法讀取圖片數據"))
		return
	}

	go func() {
		resp, err := tools.CallGeminiApiImageAndText(req.Message, imageBytes)
		resultChan <- struct {
			response string
			err      error
		}{resp, err}
	}()

	result := <-resultChan
	if result.err != nil {
		c.JSON(500, ApiResult.NewFailResult(500, result.err.Error()))
		return
	}

	c.JSON(200, ApiResult.NewSuccessResult(200, gin.H{
		"message": result.response,
	}))
}
