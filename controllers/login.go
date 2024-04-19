package controllers

import (
    "database/sql"
    "net/http"

    "golang.org/x/crypto/bcrypt"
)

type LoginController struct{}

func (lc *LoginController) Login(w http.ResponseWriter, r *http.Request) {
    session, err := store.Get(r, "session")
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    username := r.FormValue("username")
    password := r.FormValue("password")

    var dbUser string
    var dbPassword string

    err = db.QueryRow("SELECT username, password FROM users WHERE username = $1", username).Scan(&dbUser, &dbPassword)
    if err!= nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
    if err!= nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    session.Values["username"] = username
    err = session.Save(r, w)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/panel", http.StatusFound)
}