package repository

import (
	"context"
	"hacktiv8_fp_1/entity"
	"time"
)

type TodosRepository interface {
	GetTodos(ctx context.Context) (entity.Todos, error)
}

type todosConnection struct {
	filepath string
}

func NewTodosRepository(fp string) TodosRepository {
	return &todosConnection{
		filepath: fp,
	}
}

func (db *todosConnection) GetTodos(ctx context.Context) (entity.Todos, error) {
	test := entity.Todos{
		ID:          1,
		Name:        "test",
		Description: "test",
		Status:      true,
		Deadline:    time.Now(),
		BaseTimestamp: entity.BaseTimestamp{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: time.Now(),
		},
	}
	return test, nil
}
