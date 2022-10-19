package service

import (
	"context"
	"hacktiv8_fp_1/dto"
	"hacktiv8_fp_1/entity"
	"hacktiv8_fp_1/repository"
	"log"
	"strconv"

	"github.com/mashingan/smapping"
)

type TodosService interface {
	GetTodos(ctx context.Context) ([]entity.Todos, error)
	GetTodoById(ctx context.Context, id string) (entity.Todos, error)
	CreateTodo(ctx context.Context, todosDTO dto.TodosCreateDTO) (entity.Todos, error)
	UpdateTodo(ctx context.Context, id string, todoDTO dto.TodosUpdateDTO) (entity.Todos, error)
	DeleteTodoByID(ctx context.Context, id string) error
}

type todosService struct {
	todosRepository repository.TodosRepository
}

func NewTodosService(tr repository.TodosRepository) TodosService {
	return &todosService{
		todosRepository: tr,
	}
}

func (s *todosService) GetTodos(ctx context.Context) ([]entity.Todos, error) {
	todos, err := s.todosRepository.GetTodos(ctx)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return todos, nil
}

func (s *todosService) GetTodoById(ctx context.Context, id string) (entity.Todos, error) {
	todo, err := s.todosRepository.GetTodoById(ctx, id)
	if err != nil {
		log.Print(err.Error())
		return entity.Todos{}, err
	}
	return todo, nil
}

func (s *todosService) CreateTodo(ctx context.Context, todosDTO dto.TodosCreateDTO) (entity.Todos, error) {
	createdTodo := entity.Todos{}
	err := smapping.FillStruct(&createdTodo, smapping.MapFields(&todosDTO))
	if err != nil {
		return createdTodo, err
	}

	todo, err := s.todosRepository.AddNewTodoToJson(ctx, createdTodo)
	if err != nil {
		return entity.Todos{}, err
	}
	return todo, err
}

func (s *todosService) UpdateTodo(ctx context.Context, id string, todoDTO dto.TodosUpdateDTO) (entity.Todos, error) {
	var updatedTodo entity.Todos

	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return entity.Todos{}, err
	}

	updatedTodo.ID = uintId

	err = smapping.FillStruct(&updatedTodo, smapping.MapFields(&todoDTO))
	if err != nil {
		return entity.Todos{}, err
	}

	return s.todosRepository.UpdateTodo(ctx, updatedTodo)
}

func (s *todosService) DeleteTodoByID(ctx context.Context, id string) error {
	return s.todosRepository.DeleteTodo(ctx, id)
}
