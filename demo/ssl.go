package main

import (
    "fmt"
    "net/http"
    "os/exec"
)

func sslHandler(w http.ResponseWriter, r *http.Request) {
    // Get domain from request
    domain := r.FormValue("domain")

    // Create a new SSL certificate using Let's Encrypt
    cmd := exec.Command("certbot", "certonly", "--webroot", "--webroot-path=/usr/local/www/nginx/html", "--email=admin@example.com", "--agree-tos", "--non-interactive", "--expand", "--domains", "-d", domain)
    output, err := cmd.CombinedOutput()
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println(string(output))

    http.Redirect(w, r, "/", http.StatusFound)
}
