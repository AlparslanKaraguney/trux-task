package logging

import (
	"bytes"
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestLoggingInterceptor_Success(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Capture log output
	var buf bytes.Buffer
	logger.SetOutput(&buf)

	interceptor := LoggingInterceptor(logger)

	// Mock gRPC handler
	mockHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return "response", nil
	}

	// Call the interceptor
	ctx := context.Background()
	req := "test request"
	info := &grpc.UnaryServerInfo{
		FullMethod: "/test.TestService/TestMethod",
	}

	resp, err := interceptor(ctx, req, info, mockHandler)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "response", resp)

	// Verify log output
	assert.Contains(t, buf.String(), `"method":"/test.TestService/TestMethod"`)
	assert.Contains(t, buf.String(), `"request":"test request"`)
	assert.Contains(t, buf.String(), `"response":"response"`)
}

func TestLoggingInterceptor_Error(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	/// Capture log output
	var buf bytes.Buffer
	logger.SetOutput(&buf)

	interceptor := LoggingInterceptor(logger)

	// Mock gRPC handler
	mockHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, status.Errorf(codes.NotFound, "resource not found")
	}

	// Call the interceptor
	ctx := context.Background()
	req := "test request"
	info := &grpc.UnaryServerInfo{
		FullMethod: "/test.TestService/TestMethod",
	}

	resp, err := interceptor(ctx, req, info, mockHandler)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, resp)

	// Verify log output
	assert.Contains(t, buf.String(), `"method":"/test.TestService/TestMethod"`)
	assert.Contains(t, buf.String(), `"error":"resource not found"`)
	assert.Contains(t, buf.String(), `"code":`)
}
