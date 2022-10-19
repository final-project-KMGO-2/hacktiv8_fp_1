package entity

type Todos struct {
	ID          uint64 `json:"id"`
	Title       string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	IsComplete  bool   `json:"is_complete"`
	Deadline    string `json:"deadline"`
	BaseTimestamp
}
