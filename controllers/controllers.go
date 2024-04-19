package controllers

import (
    "database/sql"
    "fmt"
    "net/http"
    "text/template"

    "golang.org/x/crypto/bcrypt"

    "github.com/gorilla/sessions"
    "gorm.io/gorm"

    "webhostingpanel/models"
)

var store = sessions.NewCookieStore([]byte("secret"))

type BaseController struct{}

func (bc *BaseController) Render(w http.ResponseWriter, tmpl string, data map[string]interface{}) {
    t, _ := template.New(tmpl).ParseFiles("views/" + tmpl + ".html")
    t.Execute(w, data)
}

type LoginController struct{}

func (lc *LoginController) Login(w http.ResponseWriter, r *http.Request) {
    session, err := store.Get(r, "session")
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    username := r.FormValue("username")
    password := r.FormValue("password")

    var user models.User
    models.DB.First(&user, "username =?", username)

    if user.Username == "" {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err!= nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    session.Values["username"] = username
    err = session.Save(r, w)
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/panel", http.StatusFound)
}

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

    pc.Render(w, "panel", map[string]interface{}{"username": username})
}

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

    var domains []models.Domain
    models.DB.Where("user_id =?", username).Find(&domains)

    dc.Render(w, "domain", map[string]interface{}{"domains": domains})
}

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

    var subdomains []models.Subdomain
    models.DB.Where("user_id =?", username).Find(&subdomains)

    sc.Render(w, "subdomain", map[string]interface{}{"subdomains": subdomains})
}

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

    var ftpAccounts []models.FTPAccount
    models.DB.Where("user_id =?", username).Find(&ftpAccounts)

    fc.Render(w, "ftp", map[string]interface{}{"ftpAccounts": ftpAccounts})
}

type MySQLController struct{}

func (mc *MySQLController) MySQL(w http.ResponseWriter, r *http.Request) {
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

    var mysqlDatabases []models.MySQLDatabase
    models.DB.Where("user_id =?", username).Find(&mysqlDatabases)

    mc.Render(w, "mysql", map[string]interface{}{"mysqlDatabases": mysqlDatabases})
}

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

    var backups []models.Backup
    models.DB.Where("user_id =?", username).Find(&backups)

    bc.Render(w, "backup", map[string]interface{}{"backups": backups})
}

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

    var sslCertificates []models.SSLCertificate
    models.DB.Where("user_id =?", username).Find(&sslCertificates)

    lec.Render(w, "letsencrypt", map[string]interface{}{"sslCertificates": sslCertificates})
}

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

    var wordpressInstallations []models.WordPressInstallation
    models.DB.Where("user_id =?", username).Find(&wordpressInstallations)

    wpc.Render(w, "wordpress", map[string]interface{}{"wordpressInstallations": wordpressInstallations})
}