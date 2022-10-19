package dto

type TodosCreateDTO struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	IsComplete  bool   `json:"is_complete"`
	Deadline    string `json:"deadline"`
}

type TodosUpdateDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	IsComplete  bool   `json:"is_complete"`
	Deadline    string `json:"deadline"`
}
