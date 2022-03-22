package database

import (
	"github.com/minish144/rentateam-test/modules/config"
	"github.com/minish144/rentateam-test/modules/logger"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Module = fx.Options(
	fx.Provide(New),
)

type DatabaseModule struct {
	log    *logger.LoggerModule
	config *config.Config
	DB     *gorm.DB
}

// New creates new object of Database module
func New(log *logger.LoggerModule, config *config.Config) (*DatabaseModule, error) {
	m := &DatabaseModule{
		log:    log,
		config: config,
	}

	connectionString := m.genConnectionString()
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatalln("failed to connect to the database")
		return nil, err
	}
	m.DB = db

	return m, nil
}
