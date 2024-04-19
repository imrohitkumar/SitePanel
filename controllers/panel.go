package controllers

import (
    "net/http"
)

type PanelController struct{}

func (pc *PanelController) Panel(w http.ResponseWriter, r *http.Request) {
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

    data := map[string]interface{}{
        "username": username,
    }

    err = tmpl.ExecuteTemplate(w, "panel", data)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}