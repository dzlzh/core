package core

import (
	"fmt"

	"go.uber.org/zap"
)

func ExampleNewZap() {
	logger := NewZap("info", "logs")
	// logger.Info("Zap Example", zap.String("Test", "test"))
	sugar := logger.Sugar()
	// sugar.Info("Sugar Example")
	// zap.L().Info("Zap L Example")
	// zap.S().Info("Zap S Example")
	fmt.Printf("%T\n", logger)
	fmt.Printf("%T\n", sugar)
	fmt.Printf("%T\n", zap.L())
	fmt.Printf("%T\n", zap.S())
	// Output:
	// *zap.Logger
	// *zap.SugaredLogger
	// *zap.Logger
	// *zap.SugaredLogger
}
