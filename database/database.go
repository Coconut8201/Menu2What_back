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

	// 檢查環境變數是否完整
	if username == "" || password == "" || host == "" || port == "" || dbname == "" {
		logger.Warning("資料庫環境變數未完整設定，系統將以無資料庫模式運行")
		return createResponse(true, "系統以無資料庫模式運行")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Warning("資料庫連線失敗，系統將以無資料庫模式運行: %v", err)
		return createResponse(true, "系統以無資料庫模式運行")
	}

	// 獲取底層的 SQL DB 實例
	sqlDB, err := db.DB()
	if err != nil {
		logger.Warning("獲取資料庫實例失敗，系統將以無資料庫模式運行: %v", err)
		return createResponse(true, "系統以無資料庫模式運行")
	}

	// 設置連接池參數
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// 測試連線
	if err = sqlDB.Ping(); err != nil {
		logger.Warning("資料庫 Ping 失敗，系統將以無資料庫模式運行: %v", err)
		return createResponse(true, "系統以無資料庫模式運行")
	}

	// 在設置完資料庫連線後，加入自動遷移
	if err := db.AutoMigrate(&models.User{}); err != nil {
		logger.Warning("資料表遷移失敗，系統將以無資料庫模式運行: %v", err)
		return createResponse(true, "系統以無資料庫模式運行")
	}

	global.DB = db
	logger.Info("資料庫連線成功")
	return createResponse(true, "資料庫連線成功")
}

// GetDB 返回全局 DB 實例，增加空值檢查
func GetDB() *gorm.DB {
	if global.DB == nil {
		log.Println("警告：嘗試存取未初始化的資料庫連線")
	}
	return global.DB
}
