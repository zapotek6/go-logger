package go_logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	extLogger *zap.SugaredLogger
}

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

var logger Logger

func GlobalLogger() (*Logger, error) {
	if !logger.ExternalLoggerIsSet() {

		//rawLogger, err := zap.NewProduction()
		rawLogger, err := zap.NewDevelopment()

		if nil != err {
			return nil, err
		} else {
			logger.SetExternalLogger(rawLogger.Sugar())
		}
	}

	return &logger, nil
}
