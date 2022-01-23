package todo

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"go-todo-api/src/utils"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	Create(t *Todo) error
	GetAll() ([]Todo, error)
	GetByID(id uint) (*Todo, error)
	Update(t *Todo) error
	Delete(t *Todo) error
}

type todoRepository struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func NewTodoRepository(db *gorm.DB, r *redis.Client) *todoRepository {
	return &todoRepository{db, r}
}

func (t *todoRepository) Create(todo *Todo) error {
	go utils.ClearCache(RedisKey)
	return t.DB.Create(todo).Error
}

func (t *todoRepository) GetAll() ([]Todo, error) {
	var todos []Todo
	var context = context.Background()

	// check Todos exist in Redis cache
	result, err := t.Cache.Get(context, RedisKey).Result()

	if err != nil {
		if err := t.DB.Find(&todos).Error; err != nil {
			return nil, err
		}

		bytes, err := json.Marshal(todos)
		if err != nil {
			panic(err)
		}

		if errKey := t.Cache.Set(context, RedisKey, bytes, 30*time.Minute).Err(); errKey != nil {
			panic(errKey)
		}
	} else {
		json.Unmarshal([]byte(result), &todos)
	}

	return todos, nil
}

func (t *todoRepository) GetByID(id uint) (*Todo, error) {
	result := &Todo{}
	tx := t.DB.Where("id = ?", id).First(result)
	return result, tx.Error
}

func (t *todoRepository) Update(todo *Todo) error {
	go utils.ClearCache(RedisKey)
	return t.DB.Model(&todo).Updates(&todo).Error
}

func (t *todoRepository) Delete(todo *Todo) error {
	go utils.ClearCache(RedisKey)
	return t.DB.Delete(&todo).Error
}
