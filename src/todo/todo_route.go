package todo

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoute(app *fiber.App, db *gorm.DB, r *redis.Client) {
	api := app.Group("api")

	todoController := NewTodoController(app, NewTodoRepository(db, r))

	// Todos
	api.Get("todos", todoController.GetTodos)
	api.Post("todos", todoController.CreateTodo)
	api.Get("todos/:id", todoController.GetTodo)
	api.Put("todos/:id", todoController.UpdateTodo)
	api.Delete("todos/:id", todoController.DeleteTodo)
}
