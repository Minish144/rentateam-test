package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/minish144/rentateam-test/gen/pb"
	"github.com/minish144/rentateam-test/modules/config"
	"github.com/minish144/rentateam-test/modules/logger"
	grpcS "github.com/minish144/rentateam-test/modules/server/grpc"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var Module = fx.Options(
	fx.Provide(New),
	fx.Invoke(NewInvoke),
)

type HTTPModule struct {
	config *config.Config
	log    *logger.LoggerModule
	gwMux  *runtime.ServeMux
	mux    *http.ServeMux
}

func New(
	config *config.Config, log *logger.LoggerModule,
) (*HTTPModule, error) {
	gwMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers:  true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{},
		}),
	)
	mux := http.NewServeMux()

	mux.Handle("/", gwMux)

	m := &HTTPModule{
		config: config,
		log:    log,
		gwMux:  gwMux,
		mux:    mux,
	}

	// m.registerHandlers()

	return m, nil
}

func NewInvoke(
	lifecycle fx.Lifecycle,
	config *config.Config,
	log *logger.LoggerModule,
	grpcModule *grpcS.GRPCModule,
	httpModule *HTTPModule) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			grpcConnectionString := fmt.Sprintf("%s:%d", config.Server.GRPC.Interface, config.Server.GRPC.Port)
			httpConnectionString := fmt.Sprintf("%s:%d", config.Server.HTTP.Interface, config.Server.HTTP.Port)

			go func() {
				conn, err := grpc.DialContext(
					context.Background(),
					grpcConnectionString,
					grpc.WithBlock(),
					grpc.WithInsecure(),
				)
				if err != nil {
					log.WithFields(logrus.Fields{
						"error": err.Error(),
					}).Fatalln("failed to dial server")
				}

				err = pb.RegisterApiServiceHandler(context.Background(), httpModule.gwMux, conn)
				if err != nil {
					log.WithFields(logrus.Fields{
						"error": err.Error(),
					}).Fatalln("failed to register gateway")
				}

				if err := http.ListenAndServe(
					httpConnectionString,
					httpModule.allowCORS(httpModule.mux),
				); err != nil {
					log.WithFields(logrus.Fields{
						"error": err.Error(),
					}).Fatalln("gateway listener serve failed")
				}
			}()
			log.Infoln("serving HTTP on " + httpConnectionString)

			return nil
		},
		OnStop: func(context.Context) error {
			return nil
		},
	})
}
