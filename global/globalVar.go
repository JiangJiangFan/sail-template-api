package global

import (
	"sail-chat/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// var (
// 	Settings config.ServerConfig
// )

type Application struct {
	ConfigViper *viper.Viper // 配置文件
	Config      config.Configuration
	Log         *zap.Logger
	DB          *gorm.DB
	// Models      *Models
}

var App = new(Application)
