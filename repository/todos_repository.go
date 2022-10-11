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
	GetTodoById(ctx context.Context) (entity.Todos, error)
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

func (db *todosConnection) GetTodoById(ctx context.Context) (entity.Todos, error) {
	// ctx.Value()
	return entity.Todos{}, nil
}

func (db *todosConnection) AddNewTodoToJson(ctx context.Context) (entity.Todos, error) {
	// data, err := ioutil.ReadFile("./db/<file>.json")

	// if err != nil {
	// 	return entity.Todos{}, err
	// }

	dataMock := []entity.Todos{
		{
			ID:          69,
			Name:        "tess",
			Description: "lorem ipsum dolor sit amet maman racing",
			Status:      true,
			Deadline:    time.Now(),
			BaseTimestamp: entity.BaseTimestamp{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: time.Now(),
			},
		},
	}
	data, err := json.Marshal(dataMock)

	if err != nil {
		return entity.Todos{}, err
	}

	ioutil.WriteFile("./db/db.json", data, 0644)

	var todo []entity.Todos
	err = json.Unmarshal(data, &todo)

	if err != nil {
		return entity.Todos{}, err
	}

	return todo[0], nil

}
