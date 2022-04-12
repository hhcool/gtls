package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"

	rotate "github.com/lestrrat-go/file-rotatelogs"
)

type Format string

const (
	FormatJson    Format = "json"
	FormatConsole Format = "console"
)

var (
	logger = defaultLogger()
	sugar  = logger.Sugar()
)

type Option struct {
	ServerName string
	Format     Format
	Path       string
	MaxAge     int
	CallerSkip int
	Stacktrace zapcore.Level
}

func (o *Option) init() {
	if o.Path != "" {
		if !strings.HasSuffix(o.Path, o.ServerName) {
			o.Path = fmt.Sprintf("%s/%s", o.Path, o.ServerName)
		}
		if o.MaxAge == 0 {
			o.MaxAge = 30
		}
	}
	if o.Format == "" {
		o.Format = FormatJson
	}
}

// EnableSync
// @Description: 开启日志同步到其他
// @param option
func EnableSync(option *Option) {
	option.init()
	var cores []zapcore.Core
	cores = append(cores, consoleEncoder())

	if option.Path != "" {
		cores = append(cores, fileEncoder(option))
	}
	logger = zap.New(
		zapcore.NewTee(cores...),
	).WithOptions(
		zap.AddCaller(),
		zap.AddCallerSkip(option.CallerSkip),
		zap.AddStacktrace(option.Stacktrace),
		zap.Fields(
			zap.String("server", option.ServerName),
			zap.Namespace("detail"),
		),
	)
	sugar = logger.Sugar()
	zap.ReplaceGlobals(logger)
}

// defaultLogger
// @Description: 默认的日志实例
// @return *zap.Logger
func defaultLogger() *zap.Logger {
	cfg := baseEncoderConfig()
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zap.New(
		consoleEncoder(),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
}

// baseEncoderConfig
// @Description: 基础的配置
// @return zapcore.EncoderConfig
func baseEncoderConfig() zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()
	cfg.MessageKey = "message"
	cfg.LevelKey = "level"
	cfg.TimeKey = "time"
	cfg.NameKey = "name"
	cfg.CallerKey = "caller"
	cfg.StacktraceKey = "stack"
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	return cfg
}

// consoleEncoder
// @Description: 控制台输出编码器
// @return zapcore.Core
func consoleEncoder() zapcore.Core {
	cfg := baseEncoderConfig()
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncodeName = zapcore.FullNameEncoder
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg),
		zapcore.AddSync(zapcore.Lock(os.Stdout)),
		zapcore.DebugLevel,
	)
}

// fileEncoder
// @Description: 文件输出编码器
// @param option
// @return zapcore.Core
func fileEncoder(option *Option) zapcore.Core {
	cfg := baseEncoderConfig()
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	d, _ := time.ParseDuration(fmt.Sprintf("%dd", option.MaxAge))
	writer, _ := rotate.New(
		fmt.Sprintf("%s/%%Y%%m%%d%%H%%M.log", option.Path),
		rotate.WithLinkName(option.Path+"/latest.log"),
		rotate.WithMaxAge(d),
		rotate.WithRotationTime(24*time.Hour),
	)
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		zapcore.AddSync(writer),
		zapcore.InfoLevel,
	)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}
func With(fields ...zap.Field) *zap.Logger {
	return logger.With(fields...)
}
func WithOptions(fields ...zap.Option) *zap.Logger {
	return logger.WithOptions(fields...)
}
func Sync() {
	_ = logger.Sync()
}

func Debugf(template string, args ...interface{}) {
	sugar.Debugf(template, args...)
}
func Infof(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}
func Warnf(template string, args ...interface{}) {
	sugar.Warnf(template, args...)
}
func Errorf(template string, args ...interface{}) {
	sugar.Errorf(template, args...)
}
func Fatalf(template string, args ...interface{}) {
	sugar.Fatalf(template, args...)
}
func Panicf(template string, args ...interface{}) {
	sugar.Panicf(template, args...)
}
