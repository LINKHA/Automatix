// Code generated by goctl. DO NOT EDIT.
// Source: friend.proto

package server

import (
	"github.com/LINKHA/automatix/app/friend/cmd/rpc/internal/logic"
	"github.com/LINKHA/automatix/app/friend/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/friend/cmd/rpc/pb"
)

type RolemanagerServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedRolemanagerServer
}

func NewRolemanagerServer(svcCtx *svc.ServiceContext) *RolemanagerServer {
	return &RolemanagerServer{
		svcCtx: svcCtx,
	}
}

func (s *RolemanagerServer) AddFriend(stream pb.Rolemanager_AddFriendServer) error {
	l := logic.NewAddFriendLogic(stream.Context(), s.svcCtx)
	return l.AddFriend(stream)
}
