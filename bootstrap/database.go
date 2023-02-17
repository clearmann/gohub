package bootstrap

import (
	"errors"
	"fmt"
	"gohub/app/models/user"
	"gohub/pkg/config"
	"gohub/pkg/database"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// SetupDB 初始化数据库和 orm
func SetupDB() {
	var dbConfig gorm.Dialector
	switch config.Get[string]("database.connection") {
	case "mysql":
		{
			//构建 DSN 信息
			dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
				config.Env("database.mysql.username"),
				config.Env("database.mysql.password"),
				config.Env("database.mysql.host"),
				config.Env("database.mysql.port"),
				config.Env("database.mysql.database"),
				config.Env("database.mysql.charset"),
			)
			dbConfig = mysql.New(mysql.Config{
				DSN: dsn,
			})
		}
	case "sqlite":
		{
			//初始化
			database := config.Get[string]("database.sqlite.database")
			dbConfig = sqlite.Open(database)
		}
	default:
		panic(errors.New("database connection not supported"))
	}
	//连接数据库，并设置 GORM 的日志模式
	database.Connect(dbConfig, logger.Default.LogMode(logger.Info))
	//设置最大连接数
	database.SQLDB.SetMaxOpenConns(config.Get[int]("database.mysql.max_open_connections"))
	//设置最大空闲数
	database.SQLDB.SetMaxIdleConns(config.Get[int]("database.mysql.max_idle_connections"))
	//设置每个链接的过期数量
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.Get[int]("database.mysql.max_life_time")) * time.Second)
	database.DB.AutoMigrate(&user.User{})
}
