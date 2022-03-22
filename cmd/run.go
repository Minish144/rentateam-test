package cmd

import (
	"github.com/minish144/rentateam-test/modules"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var runServerCmd = &cobra.Command{
	Use:   "run",
	Short: "Starts Renta Team test API server",
	Long:  `Starts Renta Team test API server`,
	Run: func(cmd *cobra.Command, args []string) {
		fxApp := fx.New(modules.Module)
		fxApp.Run()
	},
}
