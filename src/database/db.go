package database

import (
	"go-todo-api/src/todo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database!")
	}
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(todo.Todo{})
}
