package Logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
	"github.com/joho/godotenv"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
)

type Logger struct {
	level  LogLevel
	logger *log.Logger
	file   *os.File
}

func init() {
	// 在套件初始化時載入 .env
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}
}

func NewLogger(level LogLevel) (*Logger, error) {
	// 從環境變數取得 logPath，如果沒有設定則使用預設值 "./logs"
	logDir := os.Getenv("logPath")
	if logDir == "" {
		logDir = "./logs"
	}

	// 確保日誌目錄存在
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("無法創建日誌目錄: %v", err)
	}

	// 生成日log
	currentTime := time.Now()
	logFileName := fmt.Sprintf("%s.log", currentTime.Format("2006-01-02"))
	logFilePath := filepath.Join(logDir, logFileName)

	// 開啟日誌檔案（追加模式）
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("無法開啟日誌檔案: %v", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)

	return &Logger{
		level:  level,
		logger: log.New(multiWriter, "", log.LstdFlags),
		file:   file,
	}, nil
}

func (l *Logger) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

// Debug 級別日誌
func (l *Logger) Debug(format string, args ...interface{}) {
	if l.level <= DEBUG {
		l.log("DEBUG", format, args...)
	}
}

// Info 級別日誌
func (l *Logger) Info(format string, args ...interface{}) {
	if l.level <= INFO {
		l.log("INFO", format, args...)
	}
}

// Warning 級別日誌
func (l *Logger) Warning(format string, args ...interface{}) {
	if l.level <= WARNING {
		l.log("WARNING", format, args...)
	}
}

// Error 級別日誌
func (l *Logger) Error(format string, args ...interface{}) {
	if l.level <= ERROR {
		l.log("ERROR", format, args...)
	}
}

// Fatal 級別日誌 - 輸出日誌後程序會終止
func (l *Logger) Fatal(format string, args ...interface{}) {
	if l.level <= FATAL {
		l.log("FATAL", format, args...)
		os.Exit(1) // 輸出日誌後終止程序
	}
}

// 內部日誌記錄方法
func (l *Logger) log(level string, format string, args ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf(format, args...)
	l.logger.Printf("[%s] %s: %s", level, timestamp, message)
}
