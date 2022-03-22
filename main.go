package main

import (
	"fmt"

	"github.com/minish144/rentateam-test/modules"
	"go.uber.org/fx"
)

func main() {
	fmt.Println("hello world")
	fxApp := fx.New(
		modules.Module,
		// fx.NopLogger, // disables FX providing logs in console
	)
	fxApp.Run()
}
