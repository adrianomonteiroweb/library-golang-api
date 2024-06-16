package database

import (
	"github.com/adrianomonteiroweb/library-golang-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() {
    database, err := gorm.Open(sqlite.Open("libraryGolang.db"), &gorm.Config{})
    if err != nil {
        panic("Error trying to connect to database: " + err.Error())
    }

    database.AutoMigrate(&models.Book{}, &models.Author{})

    DB = database
}
