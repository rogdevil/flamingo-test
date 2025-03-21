package infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

type DB struct {
	Connection *sql.DB
}

func NewDB(cfg *struct {
	Driver   string `inject:"config:db.driver"`
	Host     string `inject:"config:db.host"`
	Port     int    `inject:"config:db.port"`
	Username string `inject:"config:db.username"`
	Password string `inject:"config:db.password"`
	Database string `inject:"config:db.database"`
	SSLMode  string `inject:"config:db.sslmode"`
}) (*sql.DB, error) {
	// Load database configuration
	var dbConfig struct {
		Driver   string `inject:"config:db.driver"`
		Host     string `inject:"config:db.host"`
		Port     int    `inject:"config:db.port"`
		Username string `inject:"config:db.username"`
		Password string `inject:"config:db.password"`
		Database string `inject:"config:db.database"`
		SSLMode  string `inject:"config:db.sslmode"`
	}

	if cfg != nil {
		dbConfig.Driver = cfg.Driver
		dbConfig.Host = cfg.Host
		dbConfig.Port = cfg.Port
		dbConfig.Username = cfg.Username
		dbConfig.Password = cfg.Password
		dbConfig.Database = cfg.Database
		dbConfig.SSLMode = cfg.SSLMode
	}

	// Create PostgreSQL connection string
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Database, dbConfig.SSLMode,
	)

	// Open database connection
	db, err := sql.Open(dbConfig.Driver, dsn)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
