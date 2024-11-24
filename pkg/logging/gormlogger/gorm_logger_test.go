package gormlogger

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
)

func TestLogrusGORMLogger_Info(t *testing.T) {
	log, hook := test.NewNullLogger()
	logger := NewLogrusGORMLogger(log, logger.Info)

	// Call the Info method
	ctx := context.Background()
	logger.Info(ctx, "Info message: %s", "test")

	// Assert the log entry
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "Info message: test", hook.LastEntry().Message)
}

func TestLogrusGORMLogger_Warn(t *testing.T) {
	log, hook := test.NewNullLogger()
	logger := NewLogrusGORMLogger(log, logger.Warn)

	// Call the Warn method
	ctx := context.Background()
	logger.Warn(ctx, "Warn message: %s", "test")

	// Assert the log entry
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.WarnLevel, hook.LastEntry().Level)
	assert.Equal(t, "Warn message: test", hook.LastEntry().Message)
}

func TestLogrusGORMLogger_Error(t *testing.T) {
	log, hook := test.NewNullLogger()
	logger := NewLogrusGORMLogger(log, logger.Error)

	// Call the Error method
	ctx := context.Background()
	logger.Error(ctx, "Error message: %s", "test")

	// Assert the log entry
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "Error message: test", hook.LastEntry().Message)
}

func TestLogrusGORMLogger_Trace_Success(t *testing.T) {
	log, hook := test.NewNullLogger()
	logger := NewLogrusGORMLogger(log, logger.Info)

	// Mock the SQL execution function
	begin := time.Now()
	sqlFn := func() (string, int64) {
		return "SELECT * FROM users", 10
	}

	// Call the Trace method
	ctx := context.Background()
	logger.Trace(ctx, begin, sqlFn, nil)

	// Assert the log entry
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "SQL executed successfully", hook.LastEntry().Message)
	assert.Equal(t, "SELECT * FROM users", hook.LastEntry().Data["sql"])
	assert.Equal(t, int64(10), hook.LastEntry().Data["rows"])
}

func TestLogrusGORMLogger_Trace_Error(t *testing.T) {
	log, hook := test.NewNullLogger()
	logger := NewLogrusGORMLogger(log, logger.Info)

	// Mock the SQL execution function
	begin := time.Now()
	sqlFn := func() (string, int64) {
		return "SELECT * FROM users", 0
	}
	err := errors.New("syntax error")

	// Call the Trace method
	ctx := context.Background()
	logger.Trace(ctx, begin, sqlFn, err)

	// Assert the log entry
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "SQL execution failed", hook.LastEntry().Message)
	assert.Equal(t, "SELECT * FROM users", hook.LastEntry().Data["sql"])
	assert.Equal(t, err.Error(), hook.LastEntry().Data["error"].(error).Error())
}

func TestLogrusGORMLogger_LogMode(t *testing.T) {
	log := logrus.New()
	gromLogger := NewLogrusGORMLogger(log, logger.Warn)

	// infoLogLevel := logger.LogMode(logger.Info)

	// Change log mode
	newLogger := gromLogger.LogMode(logger.Info)

	// Assert the new logger's log level
	assert.NotEqual(t, gromLogger, newLogger)
	assert.Equal(t, logger.Info, newLogger.(*LogrusGORMLogger).level)
}

func TestLogrusGORMLogger_Silent(t *testing.T) {
	log, hook := test.NewNullLogger()
	logger := NewLogrusGORMLogger(log, logger.Silent)

	// Mock the SQL execution function
	begin := time.Now()
	sqlFn := func() (string, int64) {
		return "SELECT * FROM users", 0
	}
	err := errors.New("syntax error")

	// Call the Trace method
	ctx := context.Background()
	logger.Trace(ctx, begin, sqlFn, err)

	// Assert the log entry
	assert.Equal(t, 0, len(hook.Entries))
}
