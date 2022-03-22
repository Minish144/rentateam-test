package modules

import (
	"github.com/minish144/rentateam-test/modules/config"
	"github.com/minish144/rentateam-test/modules/database"
	"github.com/minish144/rentateam-test/modules/logger"
	"github.com/minish144/rentateam-test/modules/server/grpc"
	"github.com/minish144/rentateam-test/modules/server/http"
	"go.uber.org/fx"
)

var Module = fx.Options(
	config.Module,
	logger.Module,
	database.Module,
	grpc.Module,
	http.Module,
)
