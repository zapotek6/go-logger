package log

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strconv"
)

type Logger struct {
	extLogger *zap.SugaredLogger
}

type Level int8

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
)

func (l *Logger) SetExternalLogger(extLogger *zap.SugaredLogger) {
	l.extLogger = extLogger
}

func (l *Logger) ExternalLogger() *zap.SugaredLogger {
	return l.extLogger
}

func (l *Logger) ExternalLoggerIsSet() bool {
	return nil != l.extLogger
}

func Close() error {
	if logger.ExternalLoggerIsSet() {
		return logger.ExternalLogger().Sync()
	} else {
		return nil
	}
}

func Debug(args ...interface{}) {
	if logger.ExternalLoggerIsSet() {
		logger.ExternalLogger().Debug(args)
	}
}

func Info(args ...interface{}) {
	if logger.ExternalLoggerIsSet() {
		logger.ExternalLogger().Info(args)
	}
}

func Warn(args ...interface{}) {
	if logger.ExternalLoggerIsSet() {
		logger.ExternalLogger().Warn(args)
	}
}

func Error(args ...interface{}) {
	if logger.ExternalLoggerIsSet() {
		logger.ExternalLogger().Error(args)
	}
}

func Fatal(args ...interface{}) {
	if logger.ExternalLoggerIsSet() {
		logger.ExternalLogger().Fatal(args)
	}
}

func SetLevel(prioLevel Level) error {
	switch prioLevel {
	case DebugLevel:
		level.SetLevel(zapcore.DebugLevel)
		break
	case InfoLevel:
		level.SetLevel(zapcore.InfoLevel)
		break
	case WarnLevel:
		level.SetLevel(zapcore.WarnLevel)
		break
	case ErrorLevel:
		level.SetLevel(zapcore.ErrorLevel)
		break
	default:
		return errors.New("Invalid log level " + strconv.Itoa(int(prioLevel)))
	}

	return nil
}

var logger Logger

var level zap.AtomicLevel

func InitLogger() {
	if !logger.ExternalLoggerIsSet() {

		level = zap.NewAtomicLevel()
		level.SetLevel(zapcore.InfoLevel)

		encoderCfg := zap.NewProductionEncoderConfig()
		//encoderCfg.TimeKey = ""
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

		rawLogger := zap.New(zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.Lock(os.Stdout),
			level,
		))

		logger.SetExternalLogger(rawLogger.Sugar())
	}
}
