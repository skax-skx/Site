package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var ErrUserExists = errors.New("User already exists.")

// RegisterUser заносит пользователя по нику и хэшу пароля
// в бд, возвращает ошибку если юзер с таким ником уже есть
// в базе или при ошибке со стороны бд.
func RegisterUser(db *sql.DB, username string, passwordHash string) error {
	var err error
	_, err = db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", username, passwordHash)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrUserExists
		} else {
			return fmt.Errorf("Create user: %w", err)
		}
	} else {
		log.Println("Data saved.")
	}

	return err
}

// LoginUser возвращает boolen: true, если пользователь с таким ником
// есть и хэш пароля совпал и false, если такого пользователя нет или
// пароли не совпали, ошибку если что то пошло не так со стороны бд.
func LoginUser(db *sql.DB, username, password string) (bool, error) {
	var hash string
	err := db.QueryRow("SELECT password_hash FROM users WHERE username = $1", username).Scan(&hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	if Compare(hash, password) != nil {
		return false, nil
	}

	return true, nil
}

// Id возвращает id пользователя из бд с данным ником
// и ошибку, если такого пользователя нет или из-за
// ошибки со стороны бд.
func Id(db *sql.DB, username string) (int, error) {
	var id int
	err := db.QueryRow("SELECT id FROM users WHERE username = $1", username).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
