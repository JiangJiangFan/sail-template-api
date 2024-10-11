package bootstrap

import (
	"io"
	"log"
	"os"
	"sail-chat/global"
	"sail-chat/models"
	"strconv"
	"time"

	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	// 根据驱动进行 初始化
	switch global.App.Config.Mysql.Driver {
	case "mysql":
		return initMysql()
	default:
		return initMysql()
	}
}

func initMysql() *gorm.DB {
	dbConfig := global.App.Config.Mysql
	if dbConfig.Database == "" {
		return nil
	}
	dsn := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" + dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true, Logger: getGormLogger()}); err != nil {
		global.App.Log.Error("MySQL启动异常", zap.Any("err", err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdle)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpen)
		initMysqlTables(db)
		return db
	}
}

// 初始化数据库表
func initMysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// &global.App.Model.User{},

		// models.Config{},
		// models.Blacklist{},
		// models.Message{},
		// models.ReplyGroup{},
		// models.ReplyItem{},
		// models.Role{},
		// models.User{},
		// models.UserRole{},
		// models.Visitor{},
		// models.Welcome{},

		// models.Country{},
		// models.Abilities{},
		// models.AbilitiesUser{},
		// models.Attribute{},
		// models.Equipment{},
		// models.Incident{},
		// models.Nation{},
	)
	if err != nil {
		global.App.Log.Error("MySQL表迁移失败", zap.Any("err", err))
		os.Exit(0)
	}
	db.Model(&models.User{}).SetupJoinTable(&models.User{},"Roles",&models.UserRole{})
	
}

func getGormLogger() logger.Interface {
	var logMode logger.LogLevel
	switch global.App.Config.Mysql.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}
	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Microsecond,                       // 慢 SQL 阈值
		LogLevel:                  logMode,                                      // 日志级别
		IgnoreRecordNotFoundError: true,                                         // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  !global.App.Config.Mysql.EnableFileLogWriter, // 禁用彩色打印
	})
}

// 自定义日志
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	// 是否启用日志文件
	if global.App.Config.Mysql.EnableFileLogWriter {
		writer = &lumberjack.Logger{
			Filename:   global.App.Config.Log.Dir + "/" + global.App.Config.Mysql.LogFilename,
			MaxSize:    global.App.Config.Log.MaxSize,
			MaxBackups: global.App.Config.Log.MaxBackups,
			MaxAge:     global.App.Config.Log.MaxAge,
			Compress:   global.App.Config.Log.Compress,
		}
	} else {
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}
