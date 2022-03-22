package modules

import (
	"github.com/minish144/rentateam-test/modules/config"
	"go.uber.org/fx"
)

var Module = fx.Options(
	config.Module,
)
