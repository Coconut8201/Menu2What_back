package database

import (
	"Menu2What_back/utils/Logger"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
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

	dbConnect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		username,
		password,
		host,
		port,
		dbname,
	)

	if dbConnect == "" {
		logger.Info("資料庫連線設定錯誤：incorrect Mariadb_Url setting")
		return createResponse(false, "incorrect Mariadb_Url setting")
	}

	db, err := sql.Open("mysql", dbConnect)
	if err != nil {
		logger.Error("資料庫連線失敗: %v", err)
		return createResponse(false, fmt.Sprintf("資料庫連線失敗: %v", err))
	}

	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("關閉資料庫連線時發生錯誤: %v", err)
		}
	}()

	if err = db.Ping(); err != nil {
		logger.Error("資料庫 Ping 失敗: %v", err)
		return createResponse(false, fmt.Sprintf("資料庫 Ping 失敗: %v", err))
	}

	logger.Info("資料庫連線成功")
	return createResponse(true, "資料庫連線成功")
}
