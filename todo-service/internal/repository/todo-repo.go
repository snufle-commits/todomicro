package repository

import (
	"fmt"
	"todoservice/internal/model"

	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(todo *model.ToDo) error
	GetByID(todoID uint, userID uint) (*model.ToDo, error)
	GetAllByUser(userID uint) ([]model.ToDo, error)
	Complete(todoID uint) error
	Delete(todoID uint) error
}

type todoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) TodoRepository {
	return &todoRepo{db: db}

}

func (r *todoRepo) Create(todo *model.ToDo) error {
	if err := r.db.Create(&todo).Error; err != nil {
		return fmt.Errorf("failed to write in db")

	}
	return nil
}

func (r *todoRepo) GetByID(todoID, userID uint) (*model.ToDo, error) {
	var todo model.ToDo
	if err := r.db.Where("id = ? AND user_id = ?", todoID, userID).First(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepo) GetAllByUser(userID uint) ([]model.ToDo, error) {
	var todos []model.ToDo
	if err := r.db.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepo) Complete(todoID uint) error {
	var todo model.ToDo

	if err := r.db.First(&todo, todoID).Error; err != nil {
		return err
	}

	if err := r.db.Model(&todo).Update("completed", true).Error; err != nil {
		return err
	}

	return nil
}

func (r *todoRepo) Delete(todoID uint) error {
	if err := r.db.Unscoped().Delete(model.ToDo{}, todoID).Error; err != nil {
		return err
	}
	return nil
}
