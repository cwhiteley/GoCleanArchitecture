package appcontext

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	// Dependency of sqlx
	_ "github.com/lib/pq"
)

func getDBConnection() (*sqlx.DB, error) {
	return sqlx.Open("postgres", connectionString())
}

// TODO remove duplication for connection setting

func connectionString() string {
	userName := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	return fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", dbName, userName, password, dbHost, dbPort)
}

func DBConnectionString() string {
	userName := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", userName, password, dbHost, dbPort, dbName)
}
