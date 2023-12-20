package handler

import (
	"automatix/app/gate/cmd/api/internal/svc"
	"automatix/app/roommanager/cmd/rpc/pb"
	"context"
	"fmt"
)

func handle(serverCtx *svc.ServiceContext) {

}

func RegisterHandlers(ctx *svc.ServiceContext) {
	// //RoleManager rpc
	// createRole_Client, _ := ctx.RoleManagerRpc.CreateRoleStream(context.Background())
	// createRole_GrpcConn := svc.NewGrpcConnection(createRole_Client, "")
	// go createRole_GrpcConn.Start()

	//RoomManager rpc
	createRoleStream_Client, err := ctx.RoomManagerRpc.JoinRoomStream(context.Background())
	if err == nil {
		createRoleStream_GrpcConn := svc.NewGrpcConnection[pb.Roommanager_JoinRoomStreamClient, pb.JoinRoomStreamReq, pb.JoinRoomResp](ctx, createRoleStream_Client, "")
		go createRoleStream_GrpcConn.Start()
	} else {
		fmt.Println("rpc register err: ", err)
	}
}
