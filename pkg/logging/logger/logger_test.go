package logger

import (
	"bytes"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLogging(t *testing.T) {
	// Create a buffer to capture logs
	var buf bytes.Buffer
	mockLogger := logrus.New()
	mockLogger.SetOutput(&buf)

	// Set the custom logger
	SetLogger(mockLogger)

	// Use the logger
	log := GetLogger()
	log.Info("Test log message")

	// Verify the output
	if !bytes.Contains(buf.Bytes(), []byte("Test log message")) {
		t.Fatalf("expected log message to contain 'Test log message', got: %s", buf.String())
	}

	// Reset the logger
	SetLogger(nil)

	// Ensure the logger is re-initialized
	log = GetLogger()
	assert.NotNil(t, log)
}

func TestDevLogging(t *testing.T) {
	os.Setenv("ENV", "development")
	log := initializeLogger()

	assert.Equal(t, log.Level, logrus.DebugLevel)
}

func TestProdLogging(t *testing.T) {
	os.Setenv("ENV", "production")
	log := initializeLogger()

	assert.Equal(t, log.Level, logrus.InfoLevel)
}
