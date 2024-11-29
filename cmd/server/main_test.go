package main

import (
	"os"
	"testing"

	"github.com/AlparslanKaraguney/trux-task/internal/storage/mocks"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestValidateEnvVars(t *testing.T) {
	// Set required environment variables
	os.Setenv("DATABASE_HOST", "localhost")
	os.Setenv("DATABASE_USER", "user")
	os.Setenv("DATABASE_PASS", "pass")
	os.Setenv("DATABASE_NAME", "testdb")
	os.Setenv("DATABASE_PORT", "5432")

	// Ensure no panic occurs
	assert.NotPanics(t, func() {
		validateEnvVars([]string{
			"DATABASE_HOST", "DATABASE_USER", "DATABASE_PASS", "DATABASE_NAME", "DATABASE_PORT",
		})
	})

	// Unset a variable and check for panic
	os.Unsetenv("DATABASE_NAME")
	assert.Panics(t, func() {
		validateEnvVars([]string{
			"DATABASE_HOST", "DATABASE_USER", "DATABASE_PASS", "DATABASE_NAME", "DATABASE_PORT",
		})
	})
}

func TestSetupGRPCServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	mockStorage := mocks.NewMockStorage(ctrl)

	grpcServer, addr := setupGRPCServer(logger, mockStorage)

	assert.NotNil(t, grpcServer)
	assert.Equal(t, ":50051", addr)
}

func TestGetEnv(t *testing.T) {
	// Set an environment variable
	os.Setenv("TEST_VAR", "test")

	// Check if the variable is returned
	assert.Equal(t, "test", getEnv("TEST_VAR", "default"))

	// Check if the default value is returned
	assert.Equal(t, "default", getEnv("NOT_SET", "default"))
}
