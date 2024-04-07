package controller

import (
	"context"
	"fmt"
	"log"
	"member_service_frame/config"
	pb "member_service_frame/grpc"
	"member_service_frame/model"
	"member_service_frame/repo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedAuthorizerServer
	RedisClient repo.LoginTimeInterface
}

var _ pb.AuthorizerServer = (*Server)(nil)

func (s *Server) AuthorizeByToken(ctx context.Context, req *pb.AuthorizeToken) (*pb.AuthorizeResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	log.Println(md)

	if !ok {
		return &pb.AuthorizeResponse{Res: false}, nil
	}

	if len(md.Get("Account")) > 0 {
		var res = model.CertifyOauthAccount(s.RedisClient, ctx, md.Get("Account")[0])
		if res == model.Pass {
			return &pb.AuthorizeResponse{Res: true}, nil
		} else if res == model.LoginExpired {
			return &pb.AuthorizeResponse{Res: false}, status.Errorf(codes.PermissionDenied,
				"over 1 day from last logged in by Oauth2.0 flow , please login again.")
		} else {
			return &pb.AuthorizeResponse{Res: false}, status.Errorf(codes.PermissionDenied, "account is wrong")
		}
	}

	tokens := md.Get("Authorization")
	if tokens == nil {
		return &pb.AuthorizeResponse{Res: false}, status.Errorf(codes.PermissionDenied, "token is not found")
	}

	token := tokens[0]
	var res, reNewToken = model.CertifyToken(s.RedisClient, ctx, token)
	switch res {
	case model.Pass:
		grpc.SendHeader(ctx, metadata.Pairs("Authorization", "pass"))
		return &pb.AuthorizeResponse{Res: true}, nil
	case model.LoginExpired:
		return &pb.AuthorizeResponse{Res: false}, status.Errorf(codes.PermissionDenied,
			fmt.Sprintf("over %d days from last logged in, please login again.", config.Setting.GetValidLogin()))
	case model.TokenExpired:
		grpc.SendHeader(ctx, metadata.Pairs("Authorization", reNewToken))
		return &pb.AuthorizeResponse{Res: true}, nil
	default:
		return &pb.AuthorizeResponse{Res: false}, status.Errorf(codes.PermissionDenied, "token is wrong")
	}
}
