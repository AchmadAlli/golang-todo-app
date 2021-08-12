package database

import (
	"fmt"

	utils "github.com/AchmadAlli/golang-todo-app/app/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Connect() (*gorm.DB, error) {
	username := utils.GetEnv("DATABASE_USERNAME")
	password := utils.GetEnv("DATABASE_PASSWORD")
	database := utils.GetEnv("DATABASE_NAME")

	destination := fmt.Sprintf("%s:%s@(localhost:3306)/%s??charset=utf8mb4&parseTime=true", username, password, database)

	db, err := gorm.Open("mysql", destination)
	if err != nil {
		return nil, err
	}

	fmt.Print("Database Connected")

	db = db.Set("gorm:table_options", "DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci ENGINE=InnoDB")
	db = db.Set("gorm:auto_preload", true)

	return db, nil
}
