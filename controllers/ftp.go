package controllers

import (
    "database/sql"
    "net/http"
)

type FTPController struct{}

func (fc *FTPController) FTP(w http.ResponseWriter, r *http.Request) {
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

    var ftpAccounts []struct {
        Username string
        Password string
    }

    err = db.QueryRows("SELECT username, password FROM ftp_accounts WHERE username = $1", username).Scan(&ftpAccounts)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data := map[string]interface{}{
        "ftpAccounts": ftpAccounts,
    }

    err = tmpl.ExecuteTemplate(w, "ftp", data)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func (fc *FTPController) CreateFTP(w http.ResponseWriter, r *http.Request) {
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

    ftpUsername := r.FormValue("ftp_username")
    ftpPassword := r.FormValue("ftp_password")

    _, err = db.Exec("INSERT INTO ftp_accounts (username, password) VALUES ($1, $2)", username, ftpUsername, ftpPassword)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/ftp", http.StatusFound)
}

func (fc *FTPController) DeleteFTP(w http.ResponseWriter, r *http.Request) {
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

    ftpUsername := r.FormValue("ftp_username")

    _, err = db.Exec("DELETE FROM ftp_accounts WHERE username = $1 AND ftp_username = $2", username, ftpUsername)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/ftp", http.StatusFound)
}