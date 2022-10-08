package entity

import "time"

type Todos struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	Deadline    time.Time `json:"deadline"`
	BaseTimestamp
}
