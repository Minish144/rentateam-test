package controllers

import (
	"github.com/minish144/rentateam-test/modules/server/grpc/controllers/posts"
	"go.uber.org/fx"
)

var Module = fx.Options(
	posts.Module,
)
