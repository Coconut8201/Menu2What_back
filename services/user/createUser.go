package user

import (
	"Menu2What_back/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CreateUser 創建新使用者的服務函數
func CreateUser(db *gorm.DB, username string, password string) error {
	// 對密碼進行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 創建新使用者
	user := models.User{
		UserID:   uuid.New().String(),
		Username: username,
		Password: string(hashedPassword),
	}

	// 將使用者存入資料庫
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
