package server

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func InitDatabse(config *viper.Viper) *sql.DB {
	host := config.GetString("POSTGRES_HOST")
	port := config.GetString("POSTGRES_PORT")

	user := config.GetString("POSTGRES_USER")
	password := config.GetString("POSTGRES_PASSWORD")
	dbName := config.GetString("POSTGRES_DB_NAME")
	sslMode := config.GetString("POSTGRES_SSL")

	// Concatenate the PostgreSQL connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbName, sslMode)
	if connStr == "" {
		log.Fatalf("Database connection string is missing")
	}
	maxIdleConnections := config.GetInt("MAX_IDLE_CONNECTIONS")
	maxOpenConnections := config.GetInt("MAX_OPEN_CONNECTIONS")
	connectionMaxLifetime := config.GetDuration("CONNECTION_MAX_LIFETIME")
	driverName := config.GetString("DRIVER_NAME")
	// log.Println("PostgreSQL connection string:", connStr)
	dbHandler, err := sql.Open(driverName, connStr)
	if err != nil {
		log.Fatalf("Error while initializing database: %v", err)
	}
	dbHandler.SetMaxIdleConns(maxIdleConnections)
	dbHandler.SetMaxOpenConns(maxOpenConnections)
	dbHandler.SetConnMaxLifetime(connectionMaxLifetime)
	err = dbHandler.Ping()
	if err != nil {
		dbHandler.Close()
		log.Fatalf("Error while validating database: %v", err)
	}
	return dbHandler
}
