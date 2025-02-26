package bootstrap

import (
	"os"
	"sail-chat/global"
	"sail-chat/utils"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	level   zapcore.Level // 日志等级
	options []zap.Option  // 配置项
)

func InitializeLog() *zap.Logger {
	// 创建根目录
	createRootDir()
	// 设置日志等级
	setLogLevel()

	if global.App.Config.Log.ShowLine {
		options = append(options, zap.AddCaller())
	}
	// 初始化 zap
	return zap.New(getZapCore(), options...)

}

func createRootDir() {
	// 创建日志文件夹
	if ok, _ := utils.PathExists(global.App.Config.Log.Dir); !ok {
		// 创建日志文件夹
		_ = os.Mkdir(global.App.Config.Log.Dir, os.ModePerm)
	}
}

func setLogLevel() {
	// 设置日志等级
	switch global.App.Config.Log.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

// 扩展zap
func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder

	// 编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(global.App.Config.App.Env + `.` + l.String())
	}
	// 设置解码器
	if global.App.Config.Log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewCore(encoder, getLogWriter(), level)
}

func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   global.App.Config.Log.Dir + "/" + global.App.Config.Log.FileName,
		MaxSize:    global.App.Config.Log.MaxSize,
		MaxBackups: global.App.Config.Log.MaxBackups,
		MaxAge:     global.App.Config.Log.MaxAge,
		Compress:   global.App.Config.Log.Compress,
	}
	return zapcore.AddSync(file)
}
