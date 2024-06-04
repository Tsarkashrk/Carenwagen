package repository

import (
	"database/sql"
	"fmt"

	"github.com/Tsarkashrk/Carenwagen/server/catalog/configs"

	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg *configs.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
