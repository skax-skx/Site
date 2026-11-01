// Управление базой данных
package dataBase

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var envErr = errors.New("Env database url not found")
var db *sql.DB

// Открывает соединение с базой данных
func OpenDb() (*sql.DB, error) {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		return nil, envErr
	}

	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		return nil, envErr
	}

	db, err = sql.Open("pgx", dbURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Succesfully conected!")

	return db, err
}

// Закрывает соединение с базой данных
func CloseDb(db *sql.DB) {
	if db != nil {
		db.Close()
		log.Println("Connection was closed.")
	}
}
