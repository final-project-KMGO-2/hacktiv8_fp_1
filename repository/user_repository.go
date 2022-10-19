package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hacktiv8_fp_1/common"
	"hacktiv8_fp_1/entity"
	"io/ioutil"
	"log"
	"sort"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user entity.User) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}

type userConnection struct {
	filepath string
}

func NewUserRepository(fp string) UserRepository {
	return &userConnection{
		filepath: fp,
	}
}

func (db *userConnection) InsertUser(ctx context.Context, user entity.User) (entity.User, error) {
	byte, err := ioutil.ReadFile(db.filepath)
	if err != nil {
		return entity.User{}, err
	}
	var fileContents []entity.User
	err = json.Unmarshal(byte, &fileContents)
	if err != nil {
		return entity.User{}, err
	}
	user.Password, err = common.HashAndSalt(user.Password)
	if err != nil {
		log.Println(err.Error())
	}
	// append new user to existing list of users
	fileContents = append(fileContents, user)
	byte, err = json.Marshal(fileContents)
	if err != nil {
		log.Println(err.Error())
	}

	err = ioutil.WriteFile(db.filepath, byte, 0777)
	if err != nil {
		return entity.User{}, err
	}
	fmt.Println(db.filepath)
	return user, nil
}

func (db *userConnection) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	byte, err := ioutil.ReadFile(db.filepath)
	if err != nil {
		return entity.User{}, err
	}

	var fileContents []entity.User
	err = json.Unmarshal(byte, &fileContents)
	if err != nil {
		return entity.User{}, err
	}

	// sort slice for efficiency in searching
	sort.Slice(fileContents, func(i, j int) bool {
		return fileContents[i].Email <= fileContents[j].Email
	})

	// binary search to search email
	idx := sort.Search(len(fileContents), func(i int) bool {
		return string(fileContents[i].Email) == email
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
	if fileContents[idx-1].Email == email {
		return fileContents[idx-1], nil
	}

	return entity.User{}, errors.New("data not found")

}
