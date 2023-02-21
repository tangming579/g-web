package logger

import (
	"fmt"
	"io"
	"os"

	"go-web/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func init() {
	appName := config.AppConfig.GetString("app.name")
	mode := config.AppConfig.GetString("app.mode")
	fullPath := config.AppConfig.GetString("app.log.full.path")
	errPath := config.AppConfig.GetString("app.log.err.path")
	zapConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		TimeKey:        "log-time",
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		CallerKey:      "file",
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	logLevel := zap.DebugLevel
	if mode == "release" {
		logLevel = zap.InfoLevel
	}

	infoLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l < zapcore.WarnLevel && l >= logLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= zapcore.WarnLevel && l >= logLevel
	})
	encoder := zapcore.NewJSONEncoder(zapConfig)
	infoWriter := getWriter(fullPath, "info", appName)
	errWriter := getWriter(errPath, "error", appName)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errWriter), warnLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(zapConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), zapcore.DebugLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	Logger = logger
	defer logger.Sync()
	logger.Info("(启动中) 加载logger")

}

func Info(format string, a ...any) {
	Logger.Info(fmt.Sprintf(format, a...))
}

func Error(format string, a ...any) {
	Logger.Error(fmt.Sprintf(format, a...))
}

func Fatal(format string, a ...any) {
	Logger.Fatal(fmt.Sprintf(format, a...))
}

func Warn(format string, a ...any) {
	Logger.Warn(fmt.Sprintf(format, a...))
}

func Debug(format string, a ...any) {
	Logger.Debug(fmt.Sprintf(format, a...))
}

func InfoWithFields(msg string, fields ...zapcore.Field) {
	Logger.Info(msg, fields...)
}

func ErrorWithFields(msg string, fields ...zapcore.Field) {
	Logger.Error(msg, fields...)
}

func FatalWithFields(msg string, fields ...zapcore.Field) {
	Logger.Fatal(msg, fields...)
}

func WarnWithFields(msg string, fields ...zapcore.Field) {
	Logger.Warn(msg, fields...)
}

func DebugWithFields(msg string, fields ...zapcore.Field) {
	Logger.Debug(msg, fields...)
}

func getWriter(filePath, logType, appName string) io.Writer {
	return &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s-%s.log", filePath, appName, logType),
		MaxSize:    100,
		MaxAge:     30,
		MaxBackups: 10,
		LocalTime:  true,
		Compress:   true,
	}
}
