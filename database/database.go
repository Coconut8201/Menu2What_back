package database

import (
	"Menu2What_back/global"
	"Menu2What_back/models"
	"Menu2What_back/utils/Logger"

	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type NewConnectDbResponse struct {
	Result  bool    `json:"result"`
	Message *string `json:"message,omitempty"`
}

func createResponse(result bool, message string) NewConnectDbResponse {
	return NewConnectDbResponse{
		Result:  result,
		Message: &message,
	}
}

func NewConnectDb() NewConnectDbResponse {
	logger, err := Logger.NewLogger(Logger.INFO)
	if err != nil {
		log.Printf("無法初始化 logger: %v", err)
		return createResponse(false, "無法初始化 logger")
	}
	defer logger.Close()

	// 資料庫連線
	username := os.Getenv("Mariadb_Username")
	password := os.Getenv("Mariadb_Password")
	host := os.Getenv("Mariadb_Host")
	port := os.Getenv("Mariadb_Port")
	dbname := os.Getenv("Mariadb_Database")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbname,
	)

	if dsn == "" {
		logger.Info("資料庫連線設定錯誤：incorrect Mariadb_Url setting")
		return createResponse(false, "incorrect Mariadb_Url setting")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("資料庫連線失敗: %v", err)
		return createResponse(false, fmt.Sprintf("資料庫連線失敗: %v", err))
	}

	// 獲取底層的 SQL DB 實例
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("獲取資料庫實例失敗: %v", err)
		return createResponse(false, fmt.Sprintf("獲取資料庫實例失敗: %v", err))
	}

	// 設置連接池參數
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// 測試連線
	if err = sqlDB.Ping(); err != nil {
		logger.Error("資料庫 Ping 失敗: %v", err)
		return createResponse(false, fmt.Sprintf("資料庫 Ping 失敗: %v", err))
	}

	// 在設置完資料庫連線後，加入自動遷移
	if err := db.AutoMigrate(&models.User{}); err != nil {
		logger.Error("資料表遷移失敗: %v", err)
		return createResponse(false, fmt.Sprintf("資料表遷移失敗: %v", err))
	}

	global.DB = db
	logger.Info("資料庫連線成功")
	return createResponse(true, "資料庫連線成功")
}

// GetDB 返回全局 DB 實例
func GetDB() *gorm.DB {
	return global.DB
}
