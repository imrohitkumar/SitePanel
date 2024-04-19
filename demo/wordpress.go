package main

import (
    "fmt"
    "net/http"
    "os/exec"
)

func wordpressHandler(w http.ResponseWriter, r *http.Request) {
    // Get domain from request
    domain := r.FormValue("domain")

    // Create a new WordPress installation
    cmd := exec.Command("wp-cli", "core", "install", "--url="+domain, "--title="+domain, "--admin_user=admin", "--admin_password=password")
    output, err := cmd.CombinedOutput()
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println(string(output))

    http.Redirect(w, r, "/", http.StatusFound)
}
