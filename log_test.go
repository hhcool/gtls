package gtls

import (
	"github.com/hhcool/gtls/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestLog(t *testing.T) {
	log.EnableSync(&log.Option{
		ServerId:   "123",
		Stacktrace: zapcore.ErrorLevel,
	})
	log.Info("info", zap.String("test", "info"))
}
