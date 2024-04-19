package controllers

import (
    "database/sql"
    "net/http"
)

type BackupController struct{}

func (bc *BackupController) Backup(w http.ResponseWriter, r *http.Request) {
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

    var backups []struct {
        BackupDate string
        BackupType string
    }

    err = db.QueryRows("SELECT backup_date, backup_type FROM backups WHERE username = $1", username).Scan(&backups)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data := map[string]interface{}{
        "backups": backups,
    }

    err = tmpl.ExecuteTemplate(w, "backup", data)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}