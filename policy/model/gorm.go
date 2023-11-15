package model

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	myqlgorm "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() error {
	dialector := newDialector()
	logger := newLogger()

	gormDB, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		return errors.New("failed to connect database")
	}

	db = gormDB
	return nil
}

func newDialector() gorm.Dialector {
	return myqlgorm.Open(
		(&mysql.Config{
			Net:                  "tcp",
			User:                 os.Getenv("MYSQL_USER"),
			Passwd:               os.Getenv("MYSQL_PASSWORD"),
			Addr:                 os.Getenv("MYSQL_HOST"),
			DBName:               os.Getenv("MYSQL_DATABASE"),
			AllowNativePasswords: true,
			ParseTime:            true,
		}).FormatDSN(),
	)
}

func newLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
}
