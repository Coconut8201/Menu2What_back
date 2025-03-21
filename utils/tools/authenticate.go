package tools

import (
	"Menu2What_back/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Authenticate 驗證使用者登入
func Authenticate(db *gorm.DB, username, password string) (*models.User, error) {
	var user models.User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, errors.New("使用者不存在")
	}

	// 驗證密碼
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("密碼錯誤")
	}

	return &user, nil
}