package repositories

import (
	"github.com/AchmadAlli/golang-todo-app/app/models"
	"github.com/jinzhu/gorm"
)

type TodoRepo struct {
	DB *gorm.DB
}

func (repo *TodoRepo) Index() (*[]models.Todo, error) {
	todos := []models.Todo{}
	err := repo.DB.Model(&models.Todo{}).Find(&todos).Error
	if err != nil {
		return nil, err
	}

	return &todos, nil
}

func (repo *TodoRepo) Find(id uint) (*models.Todo, error) {
	return nil, nil
}
