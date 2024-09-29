package auth

import (
	"context"
	"encoding/base64"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var users = map[string]string{
	"Ron":  "1234",
	"Oded": "567",
}

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("--> unary interceptor: ", info.FullMethod)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	basic, ok := md["authorization"]
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "authorization header is missing")
	}

	split := strings.SplitN(basic[0], " ", 2)
	if split[0] != "Basic" {
		return nil, status.Errorf(codes.Unauthenticated, "invalid authorization header")
	}

	basicBytes, err := base64.StdEncoding.DecodeString(split[1])
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to decode authorization header")
	}

	usernameAndPassword := strings.SplitN(string(basicBytes), ":", 2)
	username, password := usernameAndPassword[0], usernameAndPassword[1]

	expectedPassword, ok := users[username]
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "username not found")
	}

	if expectedPassword != password {
		return nil, status.Errorf(codes.Unauthenticated, "incorrect password")
	}

	return handler(ctx, req)
}
