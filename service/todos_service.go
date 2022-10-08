package service

import (
	"context"
	"hacktiv8_fp_1/entity"
	"hacktiv8_fp_1/repository"
	"log"
)

type TodosService interface {
	GetTodos(ctx context.Context) (entity.Todos, error)
}

type todosService struct {
	todosRepository repository.TodosRepository
}

func NewTodosService(tr repository.TodosRepository) TodosService {
	return &todosService{
		todosRepository: tr,
	}
}

func (s *todosService) GetTodos(ctx context.Context) (entity.Todos, error) {
	todos, err := s.todosRepository.GetTodos(ctx)
	if err != nil {
		log.Print(err.Error())
		return entity.Todos{}, err
	}
	return todos, nil
}
