package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"server/global"
	"time"
)

// 日志输出常用三种方式：控制台、log文件、kafka
// 日志输出格式：默认的NewConsoleEncoder和json格式
// 日志级别

func InitZap(mode string) {
	var (
		allCore []zapcore.Core
		core    zapcore.Core
	)
	encoder := getEncoder()
	writeSyncerInfo := getLumberJackWriterInfo()
	writeSyncerError := getLumberJackWriterError()
	// 日志输出终端
	if mode == "debug" || mode == "info" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		allCore = append(allCore, zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel))
	}

	if mode == "error" {
		allCore = append(allCore, zapcore.NewCore(encoder, writeSyncerError, zapcore.ErrorLevel))
	}

	if mode == "info" {
		allCore = append(allCore, zapcore.NewCore(encoder, writeSyncerInfo, zapcore.InfoLevel))
	}

	core = zapcore.NewTee(allCore...)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()
	global.Log = logger
	global.SugarLog = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func getLumberJackWriterInfo() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./Cospk_info.log", // 日志文件位置
		MaxSize:    5,                  // 进行切割之前，日志文件最大值(单位：MB)，默认100MB
		MaxBackups: 5,                  // 保留旧文件的最大个数
		MaxAge:     1,                  // 保留旧文件的最大天数
		Compress:   false,              // 是否压缩/归档旧文件
	}

	// 输入文件和控制台
	//return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	// 只输出文件
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger))
}

func getLumberJackWriterError() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./Cospk_error.log", // 日志文件位置
		MaxSize:    5,                   // 进行切割之前，日志文件最大值(单位：MB)，默认100MB
		MaxBackups: 5,                   // 保留旧文件的最大个数
		MaxAge:     1,                   // 保留旧文件的最大天数
		Compress:   false,               // 是否压缩/归档旧文件
	}

	// 输入文件和控制台
	//return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	// 只输出文件
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger))
}
