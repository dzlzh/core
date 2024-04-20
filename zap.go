package core

import (
	"log"
	"os"
	"path"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewZap(level, directory string) *zap.Logger {
	levels := map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
	l, ok := levels[strings.ToLower(level)]
	if !ok {
		panic("not support level: " + level)
	}

	logger := zap.New(zapcore.NewTee(
		getJSONEncoderCore(l, directory),
		getConsoleEncoderCore(l),
	), zap.AddCaller())
	if l == zap.DebugLevel || l == zap.ErrorLevel {
		logger = logger.WithOptions(zap.AddStacktrace(l))
	}

	zap.ReplaceGlobals(logger)

	return logger
}

func getConsoleEncoderCore(level zapcore.Level) (core zapcore.Core) {
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(getEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		level,
	)
}

func getJSONEncoderCore(level zapcore.Level, directory string) (core zapcore.Core) {
	writer, err := writeSyncer(directory)
	if err != nil {
		log.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(getEncoderConfig()),
		writer,
		level,
	)
}

func writeSyncer(directory string) (zapcore.WriteSyncer, error) {
	writer := &lumberjack.Logger{
		Filename:   path.Join(directory, "log.log"),
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
	return zapcore.AddSync(writer), nil
	// return zapcore.NewMultiWriteSyncer(
	//     zapcore.AddSync(os.Stdout),
	//     zapcore.AddSync(writer),
	// ), nil
}

func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}