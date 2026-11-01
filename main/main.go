package main

import (
	"log"
	"myweb/dataBase"
	"myweb/handler"
	"myweb/session"
	"net/http"
)

func main() {
	db, err := dataBase.OpenDb()
	if err != nil {
		log.Fatal(err)
	}
	defer dataBase.CloseDb(db)

	http.Handle("/", http.FileServer(http.Dir("frontend")))
	http.HandleFunc("/register", handler.Register(db))
	http.HandleFunc("/login", handler.Login(db))
	http.HandleFunc("/logout", handler.Logout())

	http.HandleFunc("/home.html", func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session")
		if err != nil {
			http.Redirect(w, r, "/login.html", http.StatusSeeOther)
			return
		}

		id, ok := session.Verify(c.Value)
		if !ok {
			http.SetCookie(w, &http.Cookie{
				Name:     "session",
				Value:    "",
				MaxAge:   -1,
				Path:     "/",
				HttpOnly: true,
			})
			http.Redirect(w, r, "/login.html", http.StatusSeeOther)
			return
		}

		log.Println("user id: ", id)
		http.ServeFile(w, r, "frontend/home.html")
	})

	if err = session.Init(); err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
