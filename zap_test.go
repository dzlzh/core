package core

import (
	"fmt"
	"log/slog"

	"go.uber.org/zap"
)

func ExampleNewZap() {
	NewZap("info", "logs")
	// logger.Info("Zap Example", zap.String("Test", "test"))
	sugar := G_LOG.Sugar()
	slog.Debug("Slog Example", slog.String("Test", "test"))
	slog.Info("Slog Example", slog.String("Test", "test"))
	slog.Warn("Slog Example", slog.String("Test", "test"))
	slog.Error("Slog Example", slog.String("Test", "test"))
	// sugar.Info("Sugar Example")
	// zap.L().Info("Zap L Example")
	// zap.S().Info("Zap S Example")
	fmt.Printf("%T\n", G_LOG)
	fmt.Printf("%T\n", sugar)
	fmt.Printf("%T\n", zap.L())
	fmt.Printf("%T\n", zap.S())
	// Output:
	// *zap.Logger
	// *zap.SugaredLogger
	// *zap.Logger
	// *zap.SugaredLogger
}
