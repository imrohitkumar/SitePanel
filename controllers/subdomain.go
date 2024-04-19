package controllers

import (
    "database/sql"
    "net/http"
)

type SubdomainController struct{}

func (sc *SubdomainController) Subdomain(w http.ResponseWriter, r *http.Request) {
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

    var subdomains []string

    err = db.QueryRow("SELECT subdomain FROM subdomains WHERE username = $1", username).Scan(&subdomains)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data := map[string]interface{}{
        "subdomains": subdomains,
    }

    err = tmpl.ExecuteTemplate(w, "subdomain", data)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}