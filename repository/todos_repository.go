package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hacktiv8_fp_1/entity"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
)

type TodosRepository interface {
	GetTodos(ctx context.Context) ([]entity.Todos, error)
	GetTodoById(ctx context.Context, id string) (entity.Todos, error)
	AddNewTodoToJson(ctx context.Context, todo entity.Todos) (entity.Todos, error)
	UpdateTodo(ctx context.Context, todo entity.Todos) (entity.Todos, error)
	DeleteTodo(ctx context.Context, id string) error
}

type todosConnection struct {
	filepath string
}

func NewTodosRepository(fp string) TodosRepository {
	return &todosConnection{
		filepath: fp,
	}
}

func (db *todosConnection) GetTodos(ctx context.Context) ([]entity.Todos, error) {
	byte, err := ioutil.ReadFile(db.filepath)
	if err != nil {
		return nil, err
	}
	var fileContents []entity.Todos
	err = json.Unmarshal(byte, &fileContents)
	if err != nil {
		return nil, err
	}
	return fileContents, nil
}

func (db *todosConnection) GetTodoById(ctx context.Context, id string) (entity.Todos, error) {

	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return entity.Todos{}, err
	}
	byte, err := ioutil.ReadFile(db.filepath)
	if err != nil {
		return entity.Todos{}, err
	}

	var fileContents []entity.Todos
	err = json.Unmarshal(byte, &fileContents)
	if err != nil {
		return entity.Todos{}, err
	}

	// sort slice for efficiency in searching
	sort.Slice(fileContents, func(i, j int) bool {
		return fileContents[i].ID <= fileContents[j].ID
	})

	// binary search to search email
	idx := sort.Search(len(fileContents), func(i int) bool {
		return fileContents[i].ID >= uintId
	})

	// insert to previous file to ease next searches
	byte, err = json.Marshal(fileContents)
	if err != nil {
		log.Println(err.Error())
	}
	err = ioutil.WriteFile(db.filepath, byte, 0777)
	if err != nil {
		log.Println(err.Error())
	}

	// return data if exists
	if (idx < len(fileContents)) && fileContents[idx].ID == uintId {
		return fileContents[idx], nil
	}

	return entity.Todos{}, errors.New("data not found")
}

func (db *todosConnection) AddNewTodoToJson(ctx context.Context, todo entity.Todos) (entity.Todos, error) {

	byte, err := ioutil.ReadFile(db.filepath)
	if err != nil {
		return entity.Todos{}, err
	}

	var fileContents []entity.Todos
	err = json.Unmarshal(byte, &fileContents)
	if err != nil {
		return entity.Todos{}, err
	}

	if err != nil {
		log.Println(err.Error())
	}

	// append new user to existing list of users
	fileContents = append(fileContents, todo)
	byte, err = json.Marshal(fileContents)
	if err != nil {
		log.Println(err.Error())
	}

	err = ioutil.WriteFile(db.filepath, byte, 0777)
	if err != nil {
		return entity.Todos{}, err
	}

	return todo, nil

}

func (db *todosConnection) UpdateTodo(ctx context.Context, todo entity.Todos) (entity.Todos, error) {

	byte, err := ioutil.ReadFile(db.filepath)
	if err != nil {
		return entity.Todos{}, err
	}

	var fileContents []entity.Todos
	err = json.Unmarshal(byte, &fileContents)
	if err != nil {
		return entity.Todos{}, err
	}

	// sort slice for efficiency in searching
	sort.Slice(fileContents, func(i, j int) bool {
		return fileContents[i].ID <= fileContents[j].ID
	})

	// binary search to search id
	idx := sort.Search(len(fileContents), func(i int) bool {
		return fileContents[i].ID >= todo.ID
	})

	// update data if exists
	if (idx < len(fileContents)) && fileContents[idx].ID == todo.ID {
		fileContents[idx] = todo

		byte, err = json.Marshal(fileContents)
		if err != nil {
			log.Println(err.Error())
		}
		err = ioutil.WriteFile(db.filepath, byte, 0777)
		if err != nil {
			log.Println(err.Error())
		}

		return todo, nil
	}

	return entity.Todos{}, errors.New("data not found")
}

func (db *todosConnection) DeleteTodo(ctx context.Context, id string) error {
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	byte, err := ioutil.ReadFile(db.filepath)
	if err != nil {
		return err
	}

	var fileContents []entity.Todos
	err = json.Unmarshal(byte, &fileContents)
	if err != nil {
		return err
	}

	// sort slice for efficiency in searching
	sort.Slice(fileContents, func(i, j int) bool {
		return fileContents[i].ID <= fileContents[j].ID
	})

	// binary search to search email
	idx := sort.Search(len(fileContents), func(i int) bool {
		return fileContents[i].ID >= uintId
	})

	fmt.Println("idx : ", idx)
	// delete data if exists
	if (idx < len(fileContents)) && fileContents[idx].ID == uintId {
		fileContents[idx] = fileContents[len(fileContents)-1]

		byte, err = json.Marshal(fileContents[:len(fileContents)-1])
		if err != nil {
			log.Println(err.Error())
		}
		err = ioutil.WriteFile(db.filepath, byte, 0777)
		if err != nil {
			log.Println(err.Error())
		}

		return nil
	}

	return errors.New("data not found")
}
