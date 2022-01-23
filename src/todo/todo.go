package todo

import "go-todo-api/src/models"

type Todo struct {
	models.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        *bool  `json:"done"`
}
