package dao

import (
	"gin-ranking/api/config"
	"gin-ranking/api/pkg/logger"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Db  *gorm.DB
	err error
)

// init函数用于初始化数据库连接
func init() {
	// 使用Gorm库连接MySQL数据库
	Db, err = gorm.Open("mysql", config.Mysqldb)
	// 如果连接过程中出现错误，记录错误日志
	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect error": err.Error()})
	}
	// 如果数据库实例本身有错误，同样记录错误日志
	if Db.Error != nil {
		logger.Error(map[string]interface{}{"database error": Db.Error.Error()})
	}

	// 设置数据库连接池的最大空闲连接数为10
	Db.DB().SetMaxIdleConns(10)
	// 设置数据库连接池的最大连接数为100
	Db.DB().SetMaxOpenConns(100)
	// 设置数据库连接的最大生命周期为1小时
	Db.DB().SetConnMaxLifetime(time.Hour)
}
