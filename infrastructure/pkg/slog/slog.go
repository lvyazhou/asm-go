package slog

import (
	config2 "asm_platform/infrastructure/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var socLogger *zap.SugaredLogger

func Init() {
	appConfig := config2.GetConfig()
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	// 级别
	logLevel := appConfig.GetString("log.level")
	var l = zapcore.InfoLevel
	switch logLevel {
	case "info":
		l = zapcore.InfoLevel
	case "warn":
		l = zapcore.WarnLevel
	case "error":
		l = zapcore.ErrorLevel
	case "dpanic":
		l = zapcore.DPanicLevel
	case "panic":
		l = zapcore.PanicLevel
	case "fatal":
		l = zapcore.FatalLevel
	default:
		l = zapcore.DebugLevel
	}

	infoWriter := &lumberjack.Logger{
		Filename:   appConfig.GetString("log.file"),
		MaxSize:    appConfig.GetInt("log.max-size"),
		MaxBackups: appConfig.GetInt("log.max-backup-count"),
		MaxAge:     appConfig.GetInt("log.max-age"),
		Compress:   true,
	}

	var allCore []zapcore.Core
	allCore = append(allCore, zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), l))

	// output console
	env := appConfig.GetString("server.env")
	if env == "dev" {
		consoleDebugging := zapcore.Lock(os.Stdout)
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		allCore = append(allCore, zapcore.NewCore(consoleEncoder, consoleDebugging, l))
	}

	// 最后创建具体的Logger
	core := zapcore.NewTee(allCore...)

	// 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
	log := zap.New(core, zap.AddCaller())
	socLogger = log.Sugar()
}

func GetLogger() *zap.SugaredLogger {
	return socLogger
}

func Debug(args ...interface{}) {
	socLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	socLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	socLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	socLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	socLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	socLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	socLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	socLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	socLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	socLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	socLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	socLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	socLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	socLogger.Fatalf(template, args...)
}
