package cmd

import (
	"github.com/minish144/rentateam-test/modules/config"
	"github.com/minish144/rentateam-test/modules/database"
	"github.com/minish144/rentateam-test/modules/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrates ORM structures to db",
	Long:  `Migrates ORM structures to db`,
	Run: func(cmd *cobra.Command, args []string) {
		db := &database.DatabaseModule{}
		fx.New(
			logger.Module,
			config.Module,
			database.Module,
			fx.Populate(&db),
		)

		if err := db.Migrate(); err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Fatalln()
		} else {
			logrus.Info("migrated successfully")
		}
	},
}
