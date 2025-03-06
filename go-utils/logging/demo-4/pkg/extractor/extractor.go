package extractor

import (
	"context"
	"fmt"

	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
)

func IpExtractor(ctx context.Context) (string, error) {
	tags := grpc_ctxtags.Extract(ctx).Values()

	ip, ok := tags["ip"].(string)
	if !ok {
		return "", fmt.Errorf("ip is not extract form context")
	}

	return ip, nil
}

func RequestIDExtractor(ctx context.Context) (string, error) {
	tags := grpc_ctxtags.Extract(ctx).Values()

	rid, ok := tags["requestID"].(string)
	if !ok {
		return "", fmt.Errorf("request Id is not extract form context")
	}

	return rid, nil
}
