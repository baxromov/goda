package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitDatabase dynamically initializes the database based on the configuration
func InitDatabase(config *Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// Switch on the driver to connect to the appropriate database
	switch config.Database.Driver {
	case "postgres":
		// Compose the PostgreSQL connection string
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			config.Database.Postgres.Host,
			config.Database.Postgres.Port,
			config.Database.Postgres.User,
			config.Database.Postgres.Password,
			config.Database.Postgres.Dbname,
			config.Database.Postgres.Sslmode,
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "mysql":
		// Compose the MySQL connection string
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Database.Mysql.User,
			config.Database.Mysql.Password,
			config.Database.Mysql.Host,
			config.Database.Mysql.Port,
			config.Database.Mysql.Dbname,
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "sqlite":
		// Use the SQLite file path as the DSN
		db, err = gorm.Open(sqlite.Open(config.Database.Sqlite.Dsn), &gorm.Config{})
	default:
		log.Fatalf("Unsupported database driver: %s", config.Database.Driver)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	return db, nil
}
