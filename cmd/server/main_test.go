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

// func TestStartHTTPServer(t *testing.T) {

// 	go startHTTPServer(":8081", logrus.New()) // Start HTTP server on a test port

// 	// make sure the server has enough time to start
// 	time.Sleep(1 * time.Second)

// 	resp, err := http.Get("http://localhost:8081/health")
// 	assert.NoError(t, err)
// 	defer resp.Body.Close()

// 	body, _ := io.ReadAll(resp.Body)
// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// 	assert.Equal(t, "OK", string(body))
// }

// func TestStartGRPCServer(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	logger := logrus.New()

// 	// Mock gRPC server
// 	mockGRPC := mocksgrpc.NewMockGRPCServer(ctrl)
// 	mockGRPC.EXPECT().
// 		Serve(gomock.Any()).
// 		Return(nil)

// 	// mockGRPC.EXPECT().
// 	// 	GracefulStop().
// 	// 	Return(nil)

// 	// Use a test port
// 	testAddr := ":50052" // Use a different port to avoid conflicts

// 	startGRPCServer(mockGRPC, testAddr, logger)

// }
