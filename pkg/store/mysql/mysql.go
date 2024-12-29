package mysql

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Options MySQL配置选项
type Options struct {
	Host      string
	User      string
	Password  string
	Databasse string

	MaxIdleConns    int
	MaxOpenConns    int
	MaxConnLifeTime time.Duration
	MaxIdleTime     time.Duration

	LogLevel          int
	Logger            logger.Interface
	AutoMigrateTables []any
}

// NewMySQLClient 初始化MySQL客户端
func NewMySQLClient(opts *Options) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		opts.User, opts.Password, opts.Host, opts.Databasse,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: opts.Logger,
	})
	if err != nil {
		return nil, err
	}

	if len(opts.AutoMigrateTables) > 0 {
		err = db.AutoMigrate(opts.AutoMigrateTables...)
		if err != nil {
			return nil, err
		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(opts.MaxIdleConns)
	sqlDB.SetMaxOpenConns(opts.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(opts.MaxConnLifeTime)
	// sqlDB.SetConnMaxIdleTime(opts.MaxIdleTime)

	return db, nil
}
