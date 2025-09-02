package db

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDB() {
	var err error
	dsn := "postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable"
	DB, err = sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to PostgreSQL successfully")
}
