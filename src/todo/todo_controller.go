package todo

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

const RedisKey = "todo"

type todoController struct {
	app        *fiber.App
	repository Repository
}

func NewTodoController(app *fiber.App, repository Repository) *todoController {
	return &todoController{app, repository}
}

func (c *todoController) GetTodos(ctx *fiber.Ctx) error {
	todos, _ := c.repository.GetAll()
	return ctx.JSON(todos)
}

func (c *todoController) CreateTodo(ctx *fiber.Ctx) error {
	var todo Todo

	if err := ctx.BodyParser(&todo); err != nil {
		return err
	}

	c.repository.Create(&todo)

	return ctx.JSON(todo)
}

func (c *todoController) GetTodo(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	todo, _ := c.repository.GetByID(uint(id))

	return ctx.JSON(todo)
}

func (c *todoController) UpdateTodo(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	todo := Todo{}
	todo.Id = uint(id)

	if err := ctx.BodyParser(&todo); err != nil {
		return err
	}

	c.repository.Update(&todo)

	return ctx.JSON(todo)
}

func (c *todoController) DeleteTodo(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	todo := Todo{}
	todo.Id = uint(id)

	c.repository.Delete(&todo)

	return ctx.JSON(nil)
}
