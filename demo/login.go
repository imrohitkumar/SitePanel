package main

import (
    "database/sql"
    "fmt"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
    username := r.FormValue("username")
    password := r.FormValue("password")

    // Connect to MySQL database
    db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/panel")
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Check if username and password are correct
    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM users WHERE username =? AND password =?", username, password).Scan(&count)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if count == 0 {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    // Set session cookie
    http.SetCookie(w, &http.Cookie{
        Name:  "session",
        Value: "admin",
    })

    http.Redirect(w, r, "/", http.StatusFound)
}
