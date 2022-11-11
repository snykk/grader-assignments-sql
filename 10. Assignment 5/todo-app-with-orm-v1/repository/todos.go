package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return TodoRepository{db}
}

func (u *TodoRepository) AddTodo(todo model.Todo) error {
	return u.db.Save(&todo).Error // TODO: replace this
}

func (u *TodoRepository) ReadTodo() ([]model.Todo, error) {
	var todosFromDB []model.Todo
	err := u.db.Find(&todosFromDB).Error
	return todosFromDB, err // TODO: replace this
}

func (u *TodoRepository) UpdateDone(id uint, status bool) error {
	return u.db.Model(&model.Todo{}).Where("id = ?", id).Update("done", status).Error
}

func (u *TodoRepository) DeleteTodo(id uint) error {
	return u.db.Where("id = ?", id).Delete(&model.Todo{}).Error // TODO: replace this
}
