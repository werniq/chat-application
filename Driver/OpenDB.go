package Driver

import (
	"concurrency-chat/Logger"
	"concurrency-chat/models"
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func OpenDB() (*models.DatabaseModel, error) {
	dsn := os.Getenv("DATABASE_DSN")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		Logger.ErrorLogger().Printf("error opening database connection: %v", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		Logger.ErrorLogger().Printf("error pinging database connection: %v", err)
		return nil, err
	}
	return &models.DatabaseModel{DB: db}, nil
}
