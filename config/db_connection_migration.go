package config

import (
    "log"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    models "github.com/heronhoga/sapa-usaha-api/models"
)

var DB *gorm.DB

func DbConnectionMigration() {
    dsn := "root:@tcp(127.0.0.1:3306)/sapa_usaha?charset=utf8mb4&parseTime=True&loc=Local"

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }

    // Run the migrations
    log.Println("Running migrations...")
    err = DB.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Failed to run migrations:", err)
    }
    log.Println("Migrations completed successfully.")
	
}
