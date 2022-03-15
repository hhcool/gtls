package log

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"dataLength", "userAgent"},
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Infof("%-10s[%s]", "初始化日志库", "ok")
}

type Option struct {
	Path   string
	MaxAge int
}

func (o *Option) getPath() string {
	if o.Path == "" {
		p, err := os.Getwd()
		if err != nil {
			return "./local.log"
		}
		o.Path = fmt.Sprintf("%s/log/%s", p, "local.log")
	}
	Infof("%-10s[%s]", "日志保存路径", o.Path)
	return o.Path
}
func (o *Option) getMaxAge() time.Duration {
	if o.MaxAge == 0 {
		o.MaxAge = 30
	}
	d, _ := time.ParseDuration(fmt.Sprintf("%dd", o.MaxAge))
	Infof("%-10s[%d]", "日志保存期限", o.MaxAge)
	return d
}

func EnableFile(o Option) {
	path := o.getPath()
	writer, _ := rotate.New(
		fmt.Sprintf("%s.%%Y%%m%%d%%H%%M", path),
		rotate.WithLinkName(path),
		rotate.WithMaxAge(o.getMaxAge()),
		rotate.WithRotationTime(24*time.Hour),
	)
	pathMap := lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.PanicLevel: writer,
		logrus.ErrorLevel: writer,
		logrus.WarnLevel:  writer,
	}
	Logger.Hooks.Add(lfshook.NewHook(
		pathMap,
		&nested.Formatter{
			HideKeys:        true,
			NoColors:        true,
			NoFieldsColors:  true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
	))
}
func Info(args ...interface{}) {
	Logger.Info(args...)
}
func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}
func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}
func Error(args ...interface{}) {
	Logger.Error(args...)
}
func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}

func SafeWriterLevel(logger *logrus.Logger, level logrus.Level) *io.PipeWriter {
	return SafeEntryWriterLevel(logrus.NewEntry(logger), level)
}
func SafeEntryWriterLevel(entry *logrus.Entry, level logrus.Level) *io.PipeWriter {
	reader, writer := io.Pipe()
	var printFunc func(args ...interface{})
	switch level {
	case logrus.TraceLevel:
		printFunc = entry.Trace
	case logrus.DebugLevel:
		printFunc = entry.Debug
	case logrus.InfoLevel:
		printFunc = entry.Info
	case logrus.WarnLevel:
		printFunc = entry.Warn
	case logrus.ErrorLevel:
		printFunc = entry.Error
	case logrus.FatalLevel:
		printFunc = entry.Fatal
	case logrus.PanicLevel:
		printFunc = entry.Panic
	default:
		printFunc = entry.Print
	}
	go entryWriterScanner(entry, reader, printFunc)
	runtime.SetFinalizer(writer, writerFinalizer)
	return writer
}
func entryWriterScanner(entry *logrus.Entry, reader *io.PipeReader, printFunc func(args ...interface{})) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(scanLinesOrGiveLong)
	for scanner.Scan() {
		printFunc(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		entry.Errorf("Error while reading from Writer: %s", err)
	}
	_ = reader.Close()
}

const maxTokenLength = bufio.MaxScanTokenSize / 2

func scanLinesOrGiveLong(data []byte, atEOF bool) (advance int, token []byte, err error) {
	advance, token, err = bufio.ScanLines(data, atEOF)
	if advance > 0 || token != nil || err != nil {
		return
	}
	if len(data) < maxTokenLength {
		return
	}
	return maxTokenLength, data[0:maxTokenLength], nil
}
func writerFinalizer(writer *io.PipeWriter) {
	_ = writer.Close()
}
