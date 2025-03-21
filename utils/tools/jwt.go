package tools

import (
	"Menu2What_back/models"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("your_secret_key") // 建議使用環境變數存儲此密鑰

// GenerateJWT 生成 JWT token
func GenerateJWT(user *models.User) (string, error) {
	// 設置 token 的過期時間（例如：1小時）
	expirationTime := time.Now().Add(1 * time.Hour)

	// 創建 JWT claims
	claims := jwt.MapClaims{
		"user_id":   user.UserID,
		"username":  user.Username,
		"exp":      expirationTime.Unix(),
		"iat":      time.Now().Unix(),
	}

	// 創建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密鑰簽署 token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}