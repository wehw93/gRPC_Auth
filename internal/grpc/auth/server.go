package auth

import (
	"context"

	ssov1 "github.com/wehw93/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server){
	ssov1.RegisterAuthServer(gRPC,&serverAPI{})
}

func (s*serverAPI) Login(ctx context.Context, req *ssov1.LoginRequset) (*ssov1.LoginResponce,error){
	panic("implement me")
}

func (s*serverAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponce,error){
	panic("implement me")
}

func (s*serverAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponce,error){
	panic("implement me")
}