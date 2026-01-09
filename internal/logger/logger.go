package logger

import (
	"log"

	"go.uber.org/zap"
)

// Logger wraps zap.SugaredLogger for structured logging.
type Logger struct {
	*zap.SugaredLogger
}

// New constructs a logger; uses production config when env == "production".
func New(env string) *Logger {
	var (
		base *zap.Logger
		err  error
	)

	if env == "production" {
		base, err = zap.NewProduction()
	} else {
		base, err = zap.NewDevelopment()
	}

	if err != nil {
		log.Printf("failed to create logger: %v", err)
		base = zap.NewExample()
	}

	return &Logger{base.Sugar()}
}

func (l *Logger) Sync() {
	_ = l.SugaredLogger.Sync()
}
