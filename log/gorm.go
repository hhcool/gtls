package log

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"

	"go.uber.org/zap"
	gl "gorm.io/gorm/logger"
)

type GormLogger struct {
	ZapLogger                 *zap.Logger
	LogLevel                  gl.LogLevel
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
}
type GormOption struct {
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
}

func NewGormLogger(option *GormOption) GormLogger {
	log := GormLogger{
		ZapLogger:                 logger,
		LogLevel:                  gl.Warn,
		SlowThreshold:             option.SlowThreshold,
		SkipCallerLookup:          option.SkipCallerLookup,
		IgnoreRecordNotFoundError: option.IgnoreRecordNotFoundError,
	}
	gl.Default = log
	return log
}
func (l GormLogger) LogMode(level gl.LogLevel) gl.Interface {
	return GormLogger{
		ZapLogger:                 l.ZapLogger,
		SlowThreshold:             l.SlowThreshold,
		LogLevel:                  level,
		SkipCallerLookup:          l.SkipCallerLookup,
		IgnoreRecordNotFoundError: l.IgnoreRecordNotFoundError,
	}
}

func (l GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gl.Info {
		return
	}
	Debugf(str, args...)
}

func (l GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gl.Warn {
		return
	}
	Warnf(str, args...)
}

func (l GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gl.Error {
		return
	}
	Errorf(str, args...)
}

func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gl.Error && (!l.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		Error("trace", zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gl.Warn:
		sql, rows := fc()
		Warn("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	case l.LogLevel >= gl.Info:
		sql, rows := fc()
		Debug("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	}
}
