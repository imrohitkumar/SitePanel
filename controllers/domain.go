package controllers

import (
    "database/sql"
    "net/http"
)

type DomainController struct{}

func (dc *DomainController) Domain(w http.ResponseWriter, r *http.Request) {
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

    var domains []string

    err = db.QueryRow("SELECT domain FROM domains WHERE username = $1", username).Scan(&domains)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data := map[string]interface{}{
        "domains": domains,
    }

    err = tmpl.ExecuteTemplate(w, "domain", data)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}