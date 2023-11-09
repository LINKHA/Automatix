// Code generated by goctl. DO NOT EDIT.
// Source: servermanager.proto

package servermanager

import (
	"context"

	"automatix/app/servermanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateServerReq  = pb.CreateServerReq
	CreateServerResp = pb.CreateServerResp
	GetServerReq     = pb.GetServerReq
	GetServerResp    = pb.GetServerResp
	Server           = pb.Server
	SetServerReq     = pb.SetServerReq
	SetServerResp    = pb.SetServerResp

	Servermanager interface {
		CreateServer(ctx context.Context, in *CreateServerReq, opts ...grpc.CallOption) (*CreateServerResp, error)
		SetServer(ctx context.Context, in *SetServerReq, opts ...grpc.CallOption) (*SetServerResp, error)
		GetServer(ctx context.Context, in *GetServerReq, opts ...grpc.CallOption) (*GetServerResp, error)
	}

	defaultServermanager struct {
		cli zrpc.Client
	}
)

func NewServermanager(cli zrpc.Client) Servermanager {
	return &defaultServermanager{
		cli: cli,
	}
}

func (m *defaultServermanager) CreateServer(ctx context.Context, in *CreateServerReq, opts ...grpc.CallOption) (*CreateServerResp, error) {
	client := pb.NewServermanagerClient(m.cli.Conn())
	return client.CreateServer(ctx, in, opts...)
}

func (m *defaultServermanager) SetServer(ctx context.Context, in *SetServerReq, opts ...grpc.CallOption) (*SetServerResp, error) {
	client := pb.NewServermanagerClient(m.cli.Conn())
	return client.SetServer(ctx, in, opts...)
}

func (m *defaultServermanager) GetServer(ctx context.Context, in *GetServerReq, opts ...grpc.CallOption) (*GetServerResp, error) {
	client := pb.NewServermanagerClient(m.cli.Conn())
	return client.GetServer(ctx, in, opts...)
}
