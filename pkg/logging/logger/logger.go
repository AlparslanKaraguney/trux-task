package logger

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

var logInstance *logrus.Logger

// SetLogger allows setting a custom logger instance (useful for tests)
func SetLogger(customLogger *logrus.Logger) {
	logInstance = customLogger
}

// GetLogger returns the singleton instance of the logger
func GetLogger() *logrus.Logger {
	if logInstance == nil {
		logInstance = initializeLogger()
	}
	return logInstance
}

// initializeLogger sets up the logger instance
func initializeLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// Set log level based on the environment
	env := os.Getenv("ENV")
	if env == "development" {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	// Redirect standard log to logrus
	log.SetOutput(logger.Writer())
	log.SetFlags(0)

	return logger
}
