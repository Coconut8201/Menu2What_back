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

type GeminiAPIRequest struct {
	Message string `json:"message"`
}

func (g *GeminiController) GeminiAPI(c *gin.Context) {
	var req GeminiAPIRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error":  "無法解析請求數據",
			"detail": err.Error(),
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
		c.JSON(500, gin.H{
			"error": result.err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": result.response,
	})
}
