package database

import "fmt"

// genConnectionString returns new psql connection string
func (m *DatabaseModule) genConnectionString() string {
	host := m.config.Database.Host
	port := m.config.Database.Port
	username := m.config.Database.Username
	password := m.config.Database.Password
	db := m.config.Database.DB

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		host,
		username,
		password,
		db,
		port,
	)
}
