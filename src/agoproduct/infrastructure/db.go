package infrastructure

import (
	"database/sql"
	"fmt"

	"flamingo.me/dingo"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

type (
	// DBConfig holds the database configuration
	DBConfig struct {
		Driver   string `inject:"config:db.driver"`
		Host     string `inject:"config:db.host"`
		Port     int    `inject:"config:db.port"`
		Username string `inject:"config:db.username"`
		Password string `inject:"config:db.password"`
		Database string `inject:"config:db.database"`
		SSLMode  string `inject:"config:db.sslmode"`
	}

	// DBService provides database connectivity
	DBService struct {
		db *sql.DB
	}
)

// ConfigureDI sets up dependency injection
func ConfigureDI(injector *dingo.Injector) {
	injector.Bind(new(DBConfig)).ToProvider(NewDBConfig)
	injector.Bind(new(*sql.DB)).ToProvider(NewDB)
}

// NewDBConfig creates a new DBConfig instance
func NewDBConfig() *DBConfig {
	return &DBConfig{}
}

// NewDB creates a new database connection
func NewDB(cfg *DBConfig) (*sql.DB, error) {
	fmt.Println("Initializing database connection...", cfg)

	if cfg == nil {
		return nil, fmt.Errorf("database configuration is nil")
	}

	// Create PostgreSQL connection string
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database, cfg.SSLMode,
	)

	// Open database connection
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}
