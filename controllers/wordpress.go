package controllers

import (
    "database/sql"
    "net/http"
)

type WordPressController struct{}

func (wpc *WordPressController) WordPress(w http.ResponseWriter, r *http.Request) {
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

    var wordpressInstallations []struct {
        Domain string
        InstallationStatus string
    }

    err = db.QueryRows("SELECT domain, installation_status FROM wordpress_installations WHERE username = $1", username).Scan(&wordpressInstallations)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data := map[string]interface{}{
        "wordpressInstallations": wordpressInstallations,
    }

    err = tmpl.ExecuteTemplate(w, "wordpress", data)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}