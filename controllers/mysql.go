package controllers

import (
    "database/sql"
    "net/http"
)

type MySQLController struct{}

func (mc *MySQLController) MySQL(w http.ResponseWriter, r *http.Request) {
    session, err := store.Get(r, "session")
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    username, ok := session.Values["username"]
    if!ok {
        http.Error(w, "You are not logged in", http.StatusUnauthorized)
        return
    }

    var databases []struct {
        Database string
        Username string
        Password string
    }

    err = db.QueryRows("SELECT database, username, password FROM mysql_databases WHERE username = $1", username).Scan(&databases)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data := map[string]interface{}{
        "databases": databases,
    }

    err = tmpl.ExecuteTemplate(w, "mysql", data)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func (mc *MySQLController) CreateDatabase(w http.ResponseWriter, r *http.Request) {
    session, err := store.Get(r, "session")
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    username, ok := session.Values["username"]
    if!ok {
        http.Error(w, "You are not logged in", http.StatusUnauthorized)
        return
    }

    databaseName := r.FormValue("database_name")
    databaseUsername := r.FormValue("database_username")
    databasePassword := r.FormValue("database_password")

    _, err = db.Exec("CREATE DATABASE "+databaseName)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    _, err = db.Exec("GRANT ALL PRIVILEGES ON DATABASE "+databaseName+" TO "+databaseUsername)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    _, err = db.Exec("INSERT INTO mysql_databases (username, database, username, password) VALUES ($1, $2, $3, $4)", username, databaseName, databaseUsername, databasePassword)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/mysql", http.StatusFound)
}

func (mc *MySQLController) DeleteDatabase(w http.ResponseWriter, r *http.Request) {
    session, err := store.Get(r, "session")
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    username, ok := session.Values["username"]
    if!ok {
        http.Error(w, "You are not logged in", http.StatusUnauthorized)
        return
    }

    databaseName := r.FormValue("database_name")

    _, err = db.Exec("DROP DATABASE "+databaseName)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    _, err = db.Exec("DELETE FROM mysql_databases WHERE username = $1 AND database = $2", username, databaseName)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/mysql", http.StatusFound)
}