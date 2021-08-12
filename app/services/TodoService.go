package services

import (
	"github.com/AchmadAlli/golang-todo-app/app/models"
	"github.com/AchmadAlli/golang-todo-app/app/repositories"
)

type TodoService struct {
	repo *repositories.TodoRepo
}

func CreateService(repo *repositories.TodoRepo) *TodoService {
	return &TodoService{repo}
}

func (s *TodoService) Index() (*[]models.Todo, error) {
	todos, err := s.repo.Index()

	if err != nil {
		return nil, err
	}

	return todos, nil
}
