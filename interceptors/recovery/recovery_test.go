package recovery

import (
	"context"
	"testing"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRecoveryFunc(t *testing.T) {
	// Simulate a panic
	p := "test panic"

	// Call the recovery function
	err := RecoveryFunc(p)

	// Assertions
	assert.Error(t, err)
	st, _ := status.FromError(err)
	assert.Equal(t, codes.Internal, st.Code())
	assert.Contains(t, st.Message(), "Internal Server Error: test panic")
}

func TestRecoveryInterceptor(t *testing.T) {
	// Mock gRPC handler that panics
	mockHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
		panic("simulated panic")
	}

	// Create the recovery interceptor
	interceptor := recovery.UnaryServerInterceptor(Opts...)

	// Call the interceptor
	ctx := context.Background()
	req := "test request"
	info := &grpc.UnaryServerInfo{
		FullMethod: "/test.TestService/TestMethod",
	}

	_, err := interceptor(ctx, req, info, mockHandler)

	// Assertions
	assert.Error(t, err)
	st, _ := status.FromError(err)
	assert.Equal(t, codes.Internal, st.Code())
	assert.Contains(t, st.Message(), "Internal Server Error: simulated panic")
}
