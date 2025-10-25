package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"todoservice/internal/model"
	"todoservice/internal/repository"

	"github.com/redis/go-redis/v9"
)

type TodoService struct {
	repo  repository.TodoRepository
	redis *redis.Client
}

func NewTodoService(repo repository.TodoRepository, rdb *redis.Client) *TodoService {
	return &TodoService{repo: repo, redis: rdb}
}

func (s *TodoService) Create(todo *model.ToDo) error {
	if todo.Name == "" {
		return errors.New("Name cannot be empty")
	}
	return s.repo.Create(todo)
}

func (s *TodoService) GetByID(ctx context.Context, userID, todoID uint) (*model.ToDo, error) {
	key := fmt.Sprintf("todo:%d:%d", userID, todoID)

	val, err := s.redis.Get(ctx, key).Result()
	if err == nil {
		var todo model.ToDo
		if err := json.Unmarshal([]byte(val), &todo); err == nil {
			return &todo, nil
		}

	}

	todo, err := s.repo.GetByID(todoID, userID)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(todo)
	s.redis.Set(ctx, key, data, 2*time.Minute)

	return todo, nil
}

func (s *TodoService) GetAll(ctx context.Context, userID uint) ([]model.ToDo, error) {
	key := fmt.Sprintf("todos:%d", userID) // кэш по пользователю
	var todos []model.ToDo

	val, err := s.redis.Get(ctx, key).Result()
	if err == nil {
		json.Unmarshal([]byte(val), &todos)
		return todos, nil
	}

	todos, err = s.repo.GetAllByUser(userID)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(todos)
	s.redis.Set(ctx, key, string(data), time.Minute*5)

	return todos, nil
}

func (s *TodoService) Complete(todoID uint, userID uint) error {
	todo, err := s.repo.GetByID(todoID, userID)
	if err != nil {
		return err
	}
	if todo.Completed == true {
		return errors.New("task is already completed")
	}

	return s.repo.Complete(todoID)
}
func (s *TodoService) Delete(todoID uint) error {
	return s.repo.Delete(todoID)
}
