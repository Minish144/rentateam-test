package logger

import (
	"github.com/minish144/rentateam-test/modules/config"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// Module is fx module
var Module = fx.Provide(New)

type LoggerModule struct {
	*logrus.Logger
	config *config.Config
}

// New creates new object of this module
func New(config *config.Config) (*LoggerModule, error) {
	log := logrus.New()

	m := &LoggerModule{
		Logger: log,
		config: config,
	}

	if logLevel, err := logrus.ParseLevel(m.config.App.LogLevel); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatalln("Can't parse log level")
		return nil, err
	} else {
		m.Logger.SetLevel(logLevel)
		// m.Logger.SetReportCaller(true)
	}

	return m, nil
}
