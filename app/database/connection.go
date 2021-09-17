package database

import (
	"fmt"
	"log"

	model "github.com/AchmadAlli/golang-todo-app/app/models"
	utils "github.com/AchmadAlli/golang-todo-app/app/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Connect() (*gorm.DB, error) {
	username := utils.GetEnv("DB_USERNAME")
	password := utils.GetEnv("DB_PASSWORD")
	host := utils.GetEnv("DB_HOST")
	port := utils.GetEnv("DB_PORT")
	database := utils.GetEnv("DB_DATABASE")

	destination := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true", username, password, host, port, database)

	db, err := gorm.Open("mysql", destination)
	if err != nil {
		return nil, err
	}

	log.Println("Database Connected")

	db = db.Set("gorm:table_options", "DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci ENGINE=InnoDB")
	db = db.Set("gorm:auto_preload", true)

	return db, nil
}

func MigrateTodo(db *gorm.DB) {
	err := db.DropTableIfExists(&model.Todo{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.AutoMigrate(&model.Todo{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}
