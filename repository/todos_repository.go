package repository

import (
	"context"
	"encoding/json"
	"hacktiv8_fp_1/entity"
	"io/ioutil"
	"time"
)

type TodosRepository interface {
	GetTodos(ctx context.Context) (entity.Todos, error)
	AddNewTodoToJson(ctx context.Context) (entity.Todos, error)
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

func (db *todosConnection) AddNewTodoToJson(ctx context.Context) (entity.Todos, error) {
	data, err := ioutil.ReadFile("./db/<file>.json")
	if err != nil {
		return entity.Todos{}, err
	}

	var todo entity.Todos
	err = json.Unmarshal(data, &todo)

	if err != nil {
		return entity.Todos{}, err
	}

	return todo, nil

}
