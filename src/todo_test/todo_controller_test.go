package todo_test

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-todo-api/src/todo"
	"go-todo-api/src/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockTodoRepo struct {
	mock.Mock
}

func (m MockTodoRepo) GetAll() ([]todo.Todo, error) {
	args := m.Called()
	return args.Get(0).([]todo.Todo), args.Error(1)
}

func (m MockTodoRepo) Create(*todo.Todo) error {
	return nil
}

func (m MockTodoRepo) GetByID(id uint) (*todo.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(*todo.Todo), nil
}

func (m MockTodoRepo) Update(*todo.Todo) error {
	return nil
}

func (m MockTodoRepo) Delete(*todo.Todo) error {
	return nil
}

func TestGetTodos(t *testing.T) {
	app := fiber.New()
	todoRepoMock := new(MockTodoRepo)

	givenTodo := []todo.Todo{
		{
			Title:       "Test 1",
			Description: "",
			Done:        utils.BoolAddr(false),
		},
	}

	todoRepoMock.On("GetAll").Return(givenTodo, nil)

	todoController := todo.NewTodoController(app, todoRepoMock)
	app.Get("/api/todos", todoController.GetTodos)

	req := httptest.NewRequest("GET", "/api/todos", nil)

	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	var actualTodos []todo.Todo
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&actualTodos))
	assert.Equal(t, givenTodo, actualTodos)
}

func TestCreateTodo(t *testing.T) {
	app := fiber.New()
	todoRepoMock := new(MockTodoRepo)

	givenTodo := todo.Todo{
		Title:       "Test 1",
		Description: "",
		Done:        utils.BoolAddr(false),
	}

	todoRepoMock.On("Create").Return(givenTodo, nil)

	todoController := todo.NewTodoController(app, todoRepoMock)
	app.Post("/api/todos", todoController.CreateTodo)

	data, _ := json.Marshal(&givenTodo)
	payload := bytes.NewReader(data)
	req := httptest.NewRequest(http.MethodPost, "/api/todos", payload)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	var actualTodos todo.Todo
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&actualTodos))
	assert.Equal(t, givenTodo, actualTodos)
}

func TestGetTodo(t *testing.T) {
	app := fiber.New()
	todoRepoMock := new(MockTodoRepo)

	givenTodo := todo.Todo{
		Title:       "Test 22",
		Description: "",
		Done:        utils.BoolAddr(false),
	}
	givenTodo.Id = uint(1)

	todoRepoMock.On("GetByID", givenTodo.Id).Return(&givenTodo, nil).Once()

	todoController := todo.NewTodoController(app, todoRepoMock)
	app.Get("/api/todos/:id", todoController.GetTodo)

	req := httptest.NewRequest(http.MethodGet, "/api/todos/1", nil)

	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	var actualTodo todo.Todo
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&actualTodo))
	assert.Equal(t, givenTodo, actualTodo)
}

func TestUpdateTodo(t *testing.T) {
	app := fiber.New()
	todoRepoMock := new(MockTodoRepo)

	givenTodo := todo.Todo{
		Title:       "Test",
		Description: "",
		Done:        utils.BoolAddr(false),
	}
	givenTodo.Id = uint(1)

	todoRepoMock.On("Update", &givenTodo).Return(&givenTodo, nil).Once()

	todoController := todo.NewTodoController(app, todoRepoMock)
	app.Put("/api/todos/:id", todoController.UpdateTodo)

	data, _ := json.Marshal(&givenTodo)
	payload := bytes.NewReader(data)
	req := httptest.NewRequest(http.MethodPut, "/api/todos/1", payload)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	var actualTodo todo.Todo
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&actualTodo))
	assert.Equal(t, givenTodo, actualTodo)
}

func TestDeleteTodo(t *testing.T) {
	app := fiber.New()
	todoRepoMock := new(MockTodoRepo)

	givenTodo := todo.Todo{
		Title:       "Test",
		Description: "",
		Done:        utils.BoolAddr(false),
	}
	givenTodo.Id = uint(1)

	todoRepoMock.On("Delete", &givenTodo).Return(&givenTodo, nil).Once()

	todoController := todo.NewTodoController(app, todoRepoMock)
	app.Delete("/api/todos/:id", todoController.DeleteTodo)

	req := httptest.NewRequest(http.MethodDelete, "/api/todos/1", nil)

	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	var actualTodo todo.Todo
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&actualTodo))
	assert.Equal(t, todo.Todo{}, actualTodo)
}
