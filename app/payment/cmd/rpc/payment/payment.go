// Code generated by goctl. DO NOT EDIT!
// Source: payment.proto

package payment

import (
	"context"

	"automatix/app/payment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreatePaymentReq                     = pb.CreatePaymentReq
	CreatePaymentResp                    = pb.CreatePaymentResp
	GetPaymentBySnReq                    = pb.GetPaymentBySnReq
	GetPaymentBySnResp                   = pb.GetPaymentBySnResp
	GetPaymentSuccessRefundByOrderSnReq  = pb.GetPaymentSuccessRefundByOrderSnReq
	GetPaymentSuccessRefundByOrderSnResp = pb.GetPaymentSuccessRefundByOrderSnResp
	PaymentDetail                        = pb.PaymentDetail
	UpdateTradeStateReq                  = pb.UpdateTradeStateReq
	UpdateTradeStateResp                 = pb.UpdateTradeStateResp

	Payment interface {
		// 创建微信支付预处理订单
		CreatePayment(ctx context.Context, in *CreatePaymentReq, opts ...grpc.CallOption) (*CreatePaymentResp, error)
		// 根据sn查询流水记录
		GetPaymentBySn(ctx context.Context, in *GetPaymentBySnReq, opts ...grpc.CallOption) (*GetPaymentBySnResp, error)
		// 更新交易状态
		UpdateTradeState(ctx context.Context, in *UpdateTradeStateReq, opts ...grpc.CallOption) (*UpdateTradeStateResp, error)
		// 根据订单sn查询流水记录
		GetPaymentSuccessRefundByOrderSn(ctx context.Context, in *GetPaymentSuccessRefundByOrderSnReq, opts ...grpc.CallOption) (*GetPaymentSuccessRefundByOrderSnResp, error)
	}

	defaultPayment struct {
		cli zrpc.Client
	}
)

func NewPayment(cli zrpc.Client) Payment {
	return &defaultPayment{
		cli: cli,
	}
}

// 创建微信支付预处理订单
func (m *defaultPayment) CreatePayment(ctx context.Context, in *CreatePaymentReq, opts ...grpc.CallOption) (*CreatePaymentResp, error) {
	client := pb.NewPaymentClient(m.cli.Conn())
	return client.CreatePayment(ctx, in, opts...)
}

// 根据sn查询流水记录
func (m *defaultPayment) GetPaymentBySn(ctx context.Context, in *GetPaymentBySnReq, opts ...grpc.CallOption) (*GetPaymentBySnResp, error) {
	client := pb.NewPaymentClient(m.cli.Conn())
	return client.GetPaymentBySn(ctx, in, opts...)
}

// 更新交易状态
func (m *defaultPayment) UpdateTradeState(ctx context.Context, in *UpdateTradeStateReq, opts ...grpc.CallOption) (*UpdateTradeStateResp, error) {
	client := pb.NewPaymentClient(m.cli.Conn())
	return client.UpdateTradeState(ctx, in, opts...)
}

// 根据订单sn查询流水记录
func (m *defaultPayment) GetPaymentSuccessRefundByOrderSn(ctx context.Context, in *GetPaymentSuccessRefundByOrderSnReq, opts ...grpc.CallOption) (*GetPaymentSuccessRefundByOrderSnResp, error) {
	client := pb.NewPaymentClient(m.cli.Conn())
	return client.GetPaymentSuccessRefundByOrderSn(ctx, in, opts...)
}
