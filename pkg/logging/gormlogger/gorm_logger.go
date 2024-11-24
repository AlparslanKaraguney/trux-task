package gormlogger

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

type LogrusGORMLogger struct {
	logger *logrus.Logger
	level  logger.LogLevel
}

// NewLogrusGORMLogger creates a new instance of LogrusGORMLogger
func NewLogrusGORMLogger(log *logrus.Logger, logLevel logger.LogLevel) logger.Interface {
	return &LogrusGORMLogger{
		logger: log,
		level:  logLevel,
	}
}

// LogMode sets the logging level
func (l *LogrusGORMLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.level = level
	return &newLogger
}

// Info logs info-level messages
func (l *LogrusGORMLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Info {
		l.logger.WithContext(ctx).Infof(msg, data...)
	}
}

// Warn logs warning messages
func (l *LogrusGORMLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Warn {
		l.logger.WithContext(ctx).Warnf(msg, data...)
	}
}

// Error logs error messages
func (l *LogrusGORMLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Error {
		l.logger.WithContext(ctx).Errorf(msg, data...)
	}
}

// Trace logs SQL queries with execution time
func (l *LogrusGORMLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.level <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	entry := l.logger.WithFields(logrus.Fields{
		"sql":      sql,
		"duration": elapsed.Milliseconds(),
		"rows":     rows,
		"log_type": "sql",
	})

	if err != nil {
		entry.WithField("error", err).Error("SQL execution failed")
	} else {
		entry.Info("SQL executed successfully")
	}
}
