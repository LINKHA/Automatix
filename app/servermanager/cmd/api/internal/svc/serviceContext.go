package svc

import (
	"automatix/app/servermanager/cmd/api/internal/config"
	"automatix/app/servermanager/cmd/rpc/servermanager"

	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	ServerManagerRpc servermanager.Servermanager
	Snowflake        *snowflake.Node
	Kqueue           *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	snowflake, _ := snowflake.NewNode(c.Id)
	return &ServiceContext{
		Config:           c,
		ServerManagerRpc: servermanager.NewServermanager(zrpc.MustNewClient(c.ServerManagerRpcConf)),
		Snowflake:        snowflake,
		Kqueue:           kq.NewPusher(c.KqConf.Brokers, c.KqConf.Topic),
	}
}
