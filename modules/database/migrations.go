package database

import "github.com/minish144/rentateam-test/gen/pb"

var allModels = []interface{}{
	&pb.PostORM{},
}

func (m *DatabaseModule) Migrate() error {
	return m.DB.AutoMigrate(allModels...)
}

func (m *DatabaseModule) DropAll() error {
	// m.DB.Migrator().DropTable(allModels...)
	return m.DB.Exec("drop schema public cascade; create schema public;").Error
}
