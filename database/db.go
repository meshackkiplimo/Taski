package database

import (
    "github.com/yourusername/task-manager/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    var err error
    DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }

    // Auto-migrate the Task model
    DB.AutoMigrate(&models.Task{})
}