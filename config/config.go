package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
)

var db *gorm.DB

func ConnectDB() (*gorm.DB, error) {
    var err error
    db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
    if err!= nil {
        return nil, err
    }
    return db, nil
}