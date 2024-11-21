package recovery

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Define customfunc to handle panic
func RecoveryFunc(p any) (err error) {
	return status.Errorf(codes.Internal, "Internal Server Error: %v", p)
}

// Shared options for the logger, with a custom gRPC code to log level function.
var Opts = []recovery.Option{
	recovery.WithRecoveryHandler(RecoveryFunc),
}
