package controllers

import (
    "database/sql"
    "net/http"
)

type LetsEncryptController struct{}

func (lec *LetsEncryptController) LetsEncrypt(w http.ResponseWriter, r *http.Request) {
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

    var sslCertificates []struct {
        Domain string
        SSLStatus string
    }

    err = db.QueryRows("SELECT domain, ssl_status FROM ssl_certificates WHERE username = $1", username).Scan(&sslCertificates)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data := map[string]interface{}{
        "sslCertificates": sslCertificates,
    }

    err = tmpl.ExecuteTemplate(w, "letsencrypt", data)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}