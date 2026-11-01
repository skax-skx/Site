package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"myweb/session"
	"myweb/user"
	"net/http"
)

type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowded", http.StatusMethodNotAllowed)
			return
		}

		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if len(req.Username) < 4 {
			http.Error(w, "Username must be larger than 3", http.StatusBadRequest)
			return
		} else if len(req.Password) < 8 {
			http.Error(w, "Password must be larger than 7", http.StatusBadRequest)
			return
		}

		passwordHash, err := user.Hash(req.Password)
		if err != nil {
			log.Println("Hash error: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		err = user.RegisterUser(db, req.Username, string(passwordHash))
		if err != nil {
			if errors.Is(err, user.ErrUserExists) {
				http.Error(w, "User is exist", http.StatusConflict)
				return
			}
			log.Println("Register error: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		log.Println("User registered:", req.Username)
		w.WriteHeader(http.StatusCreated)
	}
}

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowded", http.StatusMethodNotAllowed)
			return
		}

		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		success, err := user.LoginUser(db, req.Username, req.Password)
		if err != nil {
			log.Println("Login error: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		if !success {
			log.Println("Login error: wrong username or password.")
			http.Error(w, "Wrong input", http.StatusUnauthorized)
			return
		}

		id, err := user.Id(db, req.Username)
		if err != nil {
			log.Println("User id error: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    session.Sign(id),
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: true,
		})

		log.Println("Succesfully login: ", req.Username)
		w.WriteHeader(http.StatusOK)
	}
}

func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		})
		http.Redirect(w, r, "/login.html", http.StatusSeeOther)
	}
}
