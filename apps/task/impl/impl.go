package impl

import (
	"database/sql"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/playmood/cmdb/apps/host"
	"github.com/playmood/cmdb/apps/secret"
	"github.com/playmood/cmdb/apps/task"
	"github.com/playmood/cmdb/conf"
	"google.golang.org/grpc"
)

var (
	// Service 服务实例
	svr = &impl{}
)

type impl struct {
	db  *sql.DB
	log logger.Logger
	task.UnimplementedServiceServer

	secret secret.ServiceServer
	host   host.ServiceServer
}

func (s *impl) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.log = zap.L().Named(s.Name())
	s.db = db

	// 通过mock 来解耦以来 s.secret = &secretMock{}
	s.secret = app.GetGrpcApp(secret.AppName).(secret.ServiceServer)
	s.host = app.GetGrpcApp(host.AppName).(host.ServiceServer)
	return nil
}

func (s *impl) Name() string {
	return task.AppName
}

func (s *impl) Registry(server *grpc.Server) {
	task.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
