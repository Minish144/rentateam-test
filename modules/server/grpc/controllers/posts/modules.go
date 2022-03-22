package posts

import (
	"github.com/minish144/rentateam-test/modules/config"
	"github.com/minish144/rentateam-test/modules/database"
	"github.com/minish144/rentateam-test/modules/logger"
	"go.uber.org/fx"
)

var Module = fx.Provide(New)

type GRPCControllerModule struct {
	log    *logger.LoggerModule
	config *config.Config
	db     *database.DatabaseModule
}

func New(log *logger.LoggerModule, config *config.Config, db *database.DatabaseModule) (*GRPCControllerModule, error) {
	return &GRPCControllerModule{
		log:    log,
		config: config,
		db:     db,
	}, nil
}
