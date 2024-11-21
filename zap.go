package core

import (
	"log"
	"log/slog"
	"os"
	"path"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var G_LOG *zap.Logger

func NewZap(level, directory string) {
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

	var core []zapcore.Core
	core = append(core, getJSONEncoderCore(l, directory))
	if l == zapcore.DebugLevel {
		core = append(core, getConsoleEncoderCore(l))
	}
	G_LOG = zap.New(
		zapcore.NewTee(core...),
		zap.AddCaller(),
		zap.AddStacktrace(l),
	)

	zap.ReplaceGlobals(G_LOG)

	s := slog.New(zapslog.NewHandler(
		G_LOG.Core(),
		zapslog.WithCaller(true),
		zapslog.AddStacktraceAt(slog.LevelDebug),
	))
	slog.SetDefault(s)
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
