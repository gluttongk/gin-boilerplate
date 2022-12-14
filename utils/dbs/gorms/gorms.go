package gorms

import (
	"database/sql"
	"fmt"
	"gin-boilerplate/utils/configs"
	"gin-boilerplate/utils/loggers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// InitSimpleClient generates a simple gorm client
func InitSimpleClient(usingCustomConfig bool) *gorm.DB {
	dbCfg := configs.GetGlobalDbConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", dbCfg.Usr, dbCfg.Pwd, dbCfg.Host, dbCfg.Port, dbCfg.DbName)
	//dsn := "root:pawword@tcp(127.0.0.1:3306)/test?charset=utf8mb4"
	client, err := gorm.Open(mysql.Open(dsn), getGormConfig(usingCustomConfig))

	if err != nil {
		loggers.ApiLog.Error("gorm open error: ", err)
	}

	return client

}

// InitGormClient generates a gorm client with existing database connection
func InitGormClient(customCfg bool, nativeClient *sql.DB) *gorm.DB {
	client, err := gorm.Open(mysql.New(mysql.Config{
		Conn: nativeClient,
	}), getGormConfig(customCfg))
	if err != nil {
		loggers.ApiLog.Panicln("Init gormClient failed!")
	}

	return client
}

// getGormConfig returns custom config if customCfg is true, or returns default config.
func getGormConfig(usingCustomConfig bool) *gorm.Config {
	if usingCustomConfig {
		// Refactor custom gorm configuration, here is an example
		return &gorm.Config{
			NowFunc: func() time.Time {
				//return time.Now().Local()
				return time.Now().UTC() // using UTC time
			},
			//Logger: logger.New(log.New(logs.ApiLog.Out, "\r\n", log.LUTC), logger.Config{
			//	LogLevel: logger.Error,
			//}),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			SkipDefaultTransaction: false,
		}
	} else {
		// return gorm default config
		return &gorm.Config{}
	}
}
