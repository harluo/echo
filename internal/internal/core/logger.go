package core

import (
	"fmt"
	"io"
	"os"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/labstack/echo/v4"
	labstack "github.com/labstack/gommon/log"
)

var _ echo.Logger = (*Logger)(nil)

type Logger struct {
	logger log.Logger
	prefix string
	level  labstack.Lvl
}

func newLogger(logger log.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Output() io.Writer {
	return os.Stdout
}

func (l *Logger) SetOutput(_ io.Writer) {}

func (l *Logger) Prefix() string {
	return l.prefix
}

func (l *Logger) SetPrefix(prefix string) {
	l.prefix = prefix
}

func (l *Logger) Level() (lvl labstack.Lvl) {
	switch l.logger.Level() {
	case log.LevelDebug:
		lvl = labstack.DEBUG
	case log.LevelInfo:
		lvl = labstack.INFO
	case log.LevelWarn:
		lvl = labstack.WARN
	case log.LevelError:
		lvl = labstack.ERROR
	default:
		lvl = labstack.INFO
	}

	return
}

func (l *Logger) SetLevel(lvl labstack.Lvl) {
	switch lvl {
	case labstack.DEBUG:
		l.logger.Enable(log.LevelDebug)
	case labstack.INFO:
		l.logger.Enable(log.LevelInfo)
	case labstack.WARN:
		l.logger.Enable(log.LevelWarn)
	case labstack.ERROR:
		l.logger.Enable(log.LevelError)
	default:
		l.logger.Enable(log.LevelInfo)
	}
}

func (l *Logger) SetHeader(_ string) {}

func (l *Logger) Print(fields ...any) {
	l.logger.Info(l.getMessage(), field.New("fields", fields))
}

func (l *Logger) Printf(format string, args ...any) {
	l.logger.Info(l.getMessage(), field.New("output", fmt.Sprintf(format, args...)))
}

func (l *Logger) Printj(json labstack.JSON) {
	fields := l.parseJson(json)
	l.logger.Info(l.getMessage(), fields[0], fields[1:]...)
}

func (l *Logger) Debug(fields ...any) {
	l.logger.Debug(l.getMessage(), field.New("fields", fields))
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debug(l.getMessage(), field.New("output", fmt.Sprintf(format, args...)))
}

func (l *Logger) Debugj(json labstack.JSON) {
	fields := l.parseJson(json)
	l.logger.Debug(l.getMessage(), fields[0], fields[1:]...)
}

func (l *Logger) Info(fields ...any) {
	l.logger.Info(l.getMessage(), field.New("fields", fields))
}

func (l *Logger) Infof(format string, args ...any) {
	l.logger.Info(l.getMessage(), field.New("output", fmt.Sprintf(format, args...)))
}

func (l *Logger) Infoj(json labstack.JSON) {
	fields := l.parseJson(json)
	l.logger.Info(l.getMessage(), fields[0], fields[1:]...)
}

func (l *Logger) Warn(fields ...any) {
	l.logger.Warn(l.getMessage(), field.New("fields", fields))
}

func (l *Logger) Warnf(format string, args ...any) {
	l.logger.Warn(l.getMessage(), field.New("output", fmt.Sprintf(format, args...)))
}

func (l *Logger) Warnj(json labstack.JSON) {
	fields := l.parseJson(json)
	l.logger.Warn(l.getMessage(), fields[0], fields[1:]...)
}

func (l *Logger) Error(fields ...any) {
	l.logger.Error(l.getMessage(), field.New("fields", fields))
}

func (l *Logger) Errorf(format string, args ...any) {
	l.logger.Error(l.getMessage(), field.New("output", fmt.Sprintf(format, args...)))
}

func (l *Logger) Errorj(json labstack.JSON) {
	fields := l.parseJson(json)
	l.logger.Error(l.getMessage(), fields[0], fields[1:]...)
}

func (l *Logger) Fatal(fields ...any) {
	l.logger.Fatal(l.getMessage(), field.New("fields", fields))
}

func (l *Logger) Fatalj(json labstack.JSON) {
	fields := l.parseJson(json)
	l.logger.Fatal(l.getMessage(), fields[0], fields[1:]...)
}

func (l *Logger) Fatalf(format string, args ...any) {
	l.logger.Fatal(l.getMessage(), field.New("output", fmt.Sprintf(format, args...)))
}

func (l *Logger) Panic(fields ...any) {
	l.logger.Panic(l.getMessage(), field.New("fields", fields))
}

func (l *Logger) Panicj(json labstack.JSON) {
	fields := l.parseJson(json)
	l.logger.Panic(l.getMessage(), fields[0], fields[1:]...)
}

func (l *Logger) Panicf(format string, args ...any) {
	l.logger.Panic(l.getMessage(), field.New("output", fmt.Sprintf(format, args...)))
}

func (l *Logger) getMessage() string {
	return fmt.Sprintf("%sEcho日志", l.prefix)
}

func (l *Logger) parseJson(json labstack.JSON) (fields gox.Fields[any]) {
	fields = make(gox.Fields[any], len(json))
	for key, value := range json {
		fields = append(fields, field.New(key, value))
	}

	return
}
