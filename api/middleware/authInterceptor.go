package middleware

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
)

func Unary() grpc.UnaryServerInterceptor {
	shortUrl := "/UrlServer.AntmanUrlRoutes/GetShortUrl"
	//delUrl := "/UrlServer.AntmanUrlRoutes/DeleteUrl"
	longUrl := "/UrlServer.AntmanUrlRoutes/GetLongUrl"
	signupUser := "/AntmanServer.AntmanUserRoutes/CreateNewUser"
	loginUser := "/AntmanServer.AntmanUserRoutes/LoginUser"

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Printf("--> unary interceptor: %s ", info.FullMethod)

		if info.FullMethod == loginUser || info.FullMethod == signupUser {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		//if info.FullMethod == shortUrl || info.FullMethod == delUrl || info.FullMethod == longUrl {
		if info.FullMethod == shortUrl || info.FullMethod == longUrl {
			if !ok || len(md["authorization"]) == 0 {
				md.Append("check-key", "false")
				newCtx := metadata.NewIncomingContext(ctx, md)
				return handler(newCtx, req)
			}

		}

		id, role, err := authorize(ctx)
		if err != nil {
			return fmt.Sprintf("Error : %s", err), err
		}

		if ok {
			md.Append("check-key", "true")
			md.Append("id-key", strconv.Itoa(id))
			md.Append("role-key", role)
		}
		newCtx := metadata.NewIncomingContext(ctx, md)
		return handler(newCtx, req)
	}
}

func authorize(ctx context.Context) (int, string, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return -1001, "", status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return -1002, "", status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	check, id, role, err := AuthenticateToken(accessToken)
	if err != nil || check == false {
		return -999, "", status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	if id >= 0 && role != "" {
		return id, role, nil
	}

	return -999, "", status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
