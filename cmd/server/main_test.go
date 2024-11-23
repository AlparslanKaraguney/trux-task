package main

import (
	"os"
	"testing"

	mocksGRPC "github.com/AlparslanKaraguney/trux-task/cmd/server/mocks"
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

func TestStartGRPCServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logrus.New()

	// Mock gRPC server
	mockGRPC := mocksGRPC.NewMockGRPCServer(ctrl)
	mockGRPC.EXPECT().
		Serve(gomock.Any()).
		Return(nil)

	// mockGRPC.EXPECT().
	// 	GracefulStop().
	// 	Return(nil)

	// Use a test port
	testAddr := ":50052" // Use a different port to avoid conflicts

	startGRPCServer(mockGRPC, testAddr, logger)

	// Run the gRPC server in a goroutine
	// go func() {
	// 	startGRPCServer(mockGRPC, testAddr, logger)
	// }()

	// Wait for the server to start
	// time.Sleep(1 * time.Second)

	// // Verify the port is being listened to
	// conn, err := net.Dial("tcp", "localhost:50052")
	// assert.NoError(t, err, "Expected no error when connecting to the gRPC server")
	// if conn != nil {
	// 	conn.Close()
	// }

	// Stop the gRPC server
	// mockGRPC.GracefulStop()
}

// func TestHandleGracefulShutdown(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	logger := logrus.New()

// 	// Mock gRPC server
// 	mockGRPC := mocksGRPC.NewMockGRPCServer(ctrl)
// 	mockGRPC.EXPECT().
// 		GracefulStop().
// 		Return(nil)

// 	// Mock database connection
// 	mockSQLDB := new(MockSQLDB)
// 	mockSQLDB.On("Close").Return(nil)

// 	mockGormDB := new(MockGormDB)
// 	mockGormDB.sqlDB = *mockSQLDB

// 	// Simulate exit callback
// 	exitCalled := false
// 	exitFunc := func(code int) {
// 		exitCalled = true
// 	}

// 	// Simulate signal
// 	sigChan := make(chan os.Signal, 1)
// 	signal.Notify(sigChan, os.Interrupt)
// 	go func() {
// 		time.Sleep(100 * time.Millisecond)
// 		sigChan <- os.Interrupt
// 	}()

// 	// Run the shutdown handler
// 	handleGracefulShutdown(mockGRPC, mockGormDB, logger, exitFunc)

// 	// Assertions
// 	mockGRPC.AssertCalled(t, "GracefulStop")
// 	mockSQLDB.AssertCalled(t, "Close")
// 	if !exitCalled {
// 		t.Fatalf("Expected exitFunc to be called")
// 	}
// }
