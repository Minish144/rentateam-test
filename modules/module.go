package modules

import (
	"github.com/minish144/rentateam-test/modules/config"
	"github.com/minish144/rentateam-test/modules/database"
	"github.com/minish144/rentateam-test/modules/logger"
	"go.uber.org/fx"
)

var Module = fx.Options(
	config.Module,
	logger.Module,
	database.Module,
)
