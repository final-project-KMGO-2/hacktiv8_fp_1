package seeders

import (
	"encoding/json"
	"hacktiv8_fp_1/entity"
	"io/ioutil"
	"time"
)

func SeedUser() error {
	dataMock := []entity.User{
		{
			ID:       99,
			Name:     "anam",
			Email:    "mail@anam.com",
			Password: "admin",
			BaseTimestamp: entity.BaseTimestamp{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: time.Now(),
			},
		},
	}
	data, err := json.Marshal(dataMock)

	if err != nil {
		return err
	}

	ioutil.WriteFile("./db/user.json", data, 0644)

	return nil
}
