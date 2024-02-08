package common

import (
	"context"
	"strconv"

	"google.golang.org/grpc/metadata"
)

const (
	headerAppName = "x-o3-app-name"
	headerUserID  = "x-o3-user-id"
)

// GetAppName if method is failed returns empty string
func GetAppName(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	values := md.Get(headerAppName)
	for _, v := range values {
		if v != "" {
			return v
		}
	}

	return ""
}

// GetUserID if method is failed returns zero
func GetUserID(ctx context.Context) int64 {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0
	}

	values := md.Get(headerUserID)
	for _, v := range values {
		if v != "" {
			userID, err := strconv.Atoi(v)
			if err != nil {
				return 0
			}

			return int64(userID)
		}
	}

	return 0
}
