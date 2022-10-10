package logging

import (
	"github.com/TakeoffTech/go-log/zapx"
	"go.uber.org/zap"
	"log"
)

var Logger *zap.Logger

func NewLogger(serviceName string) *zap.SugaredLogger {
	Logger, err := zapx.New(zapx.Config{
		ServiceName: serviceName,
	})

	if err != nil {
		log.Printf(`{"severity": "error", "message": "failed to initialize zap logger: %v"}`, err)

		Logger = zap.S()
	}

	return Logger
}
