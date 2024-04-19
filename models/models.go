package models

import (
    "gorm.io/gorm"
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

func init() {
    models.DB.AutoMigrate(&models.User{}, &models.Domain{}, &models.Subdomain{}, &models.FTPAccount{}, &models.MySQLDatabase{}, &models.Backup{}, &models.SSLCertificate{}, &models.WordPressInstallation{})
}