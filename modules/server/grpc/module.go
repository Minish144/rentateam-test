package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/minish144/rentateam-test/gen/pb"
	"github.com/minish144/rentateam-test/modules/config"
	"github.com/minish144/rentateam-test/modules/logger"
	"github.com/minish144/rentateam-test/modules/server/grpc/controllers"
	"github.com/minish144/rentateam-test/modules/server/grpc/controllers/posts"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type GRPCModule struct {
	config *config.Config
	log    *logger.LoggerModule
	server *grpc.Server
}

var Module = fx.Options(
	controllers.Module,
	fx.Provide(New),
	fx.Invoke(NewInvoke),
)

func New(
	log *logger.LoggerModule,
	config *config.Config,
	postsCtrl *posts.GRPCControllerModule,
) (*GRPCModule, error) {
	s := &ApiServiceServer{
		postsCtrl: postsCtrl,
	}
	grpcServer := grpc.NewServer()

	pb.RegisterApiServiceServer(grpcServer, s)

	m := &GRPCModule{
		log:    log,
		config: config,
		server: grpcServer,
	}

	return m, nil
}

func NewInvoke(
	lifecycle fx.Lifecycle,
	log *logger.LoggerModule,
	config *config.Config,
	grpcModule *GRPCModule) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			grpcConnectionString := fmt.Sprintf("%s:%d", config.Server.GRPC.Interface, config.Server.GRPC.Port)
			log.Infoln("serving gRPC on " + grpcConnectionString)
			go func() {
				lis, err := net.Listen("tcp", grpcConnectionString)
				if err != nil {
					log.WithFields(logrus.Fields{
						"error":           err.Error(),
						"conectionString": grpcConnectionString,
					}).Fatalln("failed to listen")
				}

				if err := grpcModule.server.Serve(lis); err != nil {
					log.WithFields(logrus.Fields{
						"error": err.Error(),
					}).Fatalln("failed to start server")
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			grpcModule.server.Stop()
			return nil
		},
	})
}
