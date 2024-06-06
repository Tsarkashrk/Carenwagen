package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Tsarkashrk/Carenwagen/server/auth/config"
	_ "github.com/lib/pq"
)

func InitDB(config *config.Config) (*sql.DB, error) {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
		return nil, err
	}
	return db, nil
}
