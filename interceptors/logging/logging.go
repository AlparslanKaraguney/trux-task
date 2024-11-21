package logging

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// LoggingInterceptor logs the details of each gRPC call. UnaryServerInterceptor
func LoggingInterceptor(logger *logrus.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()

		// Extract metadata from context
		md, _ := metadata.FromIncomingContext(ctx)
		fields := logrus.Fields{
			"method":     info.FullMethod,
			"metadata":   md,
			"request":    req,
			"start_time": start.Format(time.RFC3339),
			"log_type":   "request",
		}

		// Call the handler
		resp, err := handler(ctx, req)

		// Add response and duration
		fields["response"] = resp
		fields["duration_ms"] = time.Since(start).Milliseconds()

		if err != nil {
			st := status.Convert(err)
			fields["error"] = st.Message()
			fields["code"] = st.Code()
			logger.WithFields(fields).Error("gRPC request failed")
		} else {
			logger.WithFields(fields).Info("gRPC request completed")
		}

		return resp, err
	}
}
