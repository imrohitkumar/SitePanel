package main

import (
    "database/sql"
    "fmt"
    "html/template"
    "net/http"

    "golang.org/x/crypto/bcrypt"

    "github.com/gorilla/sessions"
    "gorm.io/gorm"

    "webhostingpanel/config"
    "webhostingpanel/controllers"
    "webhostingpanel/models"
)

var (
    db  *gorm.DB
    tmpl *template.Template
    store = sessions.NewCookieStore([]byte("secret"))
)

func main() {
    var err error
    db, err = config.ConnectDB()
    if err!= nil {
        fmt.Println(err)
        return
    }

    tmpl = template.Must(template.New("").Funcs(template.FuncMap{}).ParseGlob("views/*.html"))

    http.HandleFunc("/", index)
    http.HandleFunc("/login", login)
    http.HandleFunc("/panel", panel)
    http.HandleFunc("/domain", domain)
    http.HandleFunc("/subdomain", subdomain)
    http.HandleFunc("/ftp", ftp)
    http.HandleFunc("/mysql", mysql)
    http.HandleFunc("/backup", backup)
    http.HandleFunc("/letsencrypt", letsencrypt)
    http.HandleFunc("/wordpress", wordpress)

    fmt.Println("Server started on port 8080")
    http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/login", http.StatusFound)
}

func login(w http.ResponseWriter, r *http.Request) {
    ctrl := controllers.LoginController{}
    ctrl.Login(w, r)
}

func panel(w http.ResponseWriter, r *http.Request) {
    ctrl := controllers.PanelController{}
    ctrl.Panel(w, r)
}

func domain(w http.ResponseWriter, r *http.Request) {
    ctrl := controllers.DomainController{}
    ctrl.Domain(w, r)
}

func subdomain(w http.ResponseWriter, r *http.Request) {
    ctrl := controllers.SubdomainController{}
    ctrl.Subdomain(w, r)
}

func ftp(w http.ResponseWriter, r *http.Request) {
    ctrl := controllers.FTPController{}
    ctrl.FTP(w, r)
}

func mysql(w http.ResponseWriter, r *http.Request) {
    ctrl := controllers.MySQLController{}
    ctrl.MySQL(w, r)
}

func backup(w http.ResponseWriter, r *http.Request) {
    ctrl := controllers.BackupController{}
    ctrl.Backup(w, r)
}

func letsencrypt(w http.ResponseWriter, r *http.Request) {
    ctrl := controllers.LetsEncryptController{}
    ctrl.LetsEncrypt(w, r)
}

func wordpress(w http.ResponseWriter, r *http.Request) {
    ctrl := controllers.WordPressController{}
    ctrl.WordPress(w, r)
}