package main

import (
    "database/sql"
    "encoding/gob"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "os"
    "path/filepath"

    "github.com/gorilla/sessions"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var (
    db  *gorm.DB
    err error
    store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
    tmpl *template.Template
)

type User struct {
    gorm.Model
    Username string `gorm:"uniqueIndex"`
    Password string
}

type Domain struct {
    gorm.Model
    UserID  uint
    User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    Domain string
}

type Subdomain struct {
    gorm.Model
    UserID  uint
    User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    Subdomain string
}

type FTPAccount struct {
    gorm.Model
    UserID  uint
    User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    Username string
    Password string
}

type MySQLDatabase struct {
    gorm.Model
    UserID  uint
    User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    Database string
    Username string
    Password string
}

type Backup struct {
    gorm.Model
    UserID  uint
    User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    BackupDate string
    BackupType string
}

type SSLCertificate struct {
    gorm.Model
    UserID  uint
    User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    Domain string
    SSLStatus string
}

type WordPressInstallation struct {
    gorm.Model
    UserID  uint
    User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    Domain string
    InstallationStatus string
}

func main() {
    // Initialize database connection
    dsn := os.Getenv("DATABASE_URL")
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err!= nil {
        log.Fatal(err)
    }

    // Migrate database schema
    db.AutoMigrate(&User{}, &Domain{}, &Subdomain{}, &FTPAccount{}, &MySQLDatabase{}, &Backup{}, &SSLCertificate{}, &WordPressInstallation{})

    // Initialize template engine
    tmpl = template.Must(template.New("").Funcs(template.FuncMap{}).ParseGlob("views/*.html"))

    // Initialize session store
    gob.Register(map[string]interface{}{})

    // Define routes
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

    // Start server
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