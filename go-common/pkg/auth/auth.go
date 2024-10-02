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

type UserId struct{}

type User struct {
	Username string
	Password string
}

var users = map[string]User{
	"003ba3fa-9434-44bf-a9a2-e982659e40ec": {
		Username: "Oded",
		Password: "f54c3889001199907a54d76975639f18",
	},
	"7a52a58b-5aa9-4d2f-a05c-c0d090bb1b2a": {
		Username: "Eylon",
		Password: "8fa3fbef1b6518a80d6b9be774e8dc58",
	},
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

	var found = false

	for uuid, user := range users {
		if user.Username == username && user.Password == password {
			found = true
			ctx = context.WithValue(ctx, UserId{}, uuid)
			break
		}
	}

	if !found {
		return nil, status.Errorf(codes.Unauthenticated, "incorrect username or password")
	}

	return handler(ctx, req)
}
