package service

import (
	"github.com/sergey-timtsunyk/todo/pkg/data"
	"github.com/sergey-timtsunyk/todo/pkg/repository"
)

type TodoListService struct {
	repository repository.TodoList
}

func NewTodoListService(repository repository.TodoList) *TodoListService {
	return &TodoListService{repository: repository}
}

func (t *TodoListService) Create(userId uint, list data.List) (uint, error) {
	return t.repository.Create(userId, list)
}

func (t *TodoListService) GetAll(userId uint) ([]data.List, error) {
	return t.repository.GetAll(userId)
}

func (t *TodoListService) GetByIdAndUserId(userId uint, listId uint) (data.List, error) {
	return t.repository.GetByIdAndUserId(userId, listId)
}

func (t *TodoListService) DeleteByIdAndUserId(userId uint, listId uint) error {
	return t.repository.DeleteByIdAndUserId(userId, listId)
}

func (t *TodoListService) UpdateList(userId uint, listId uint, updateList data.UpdateListInput) error {
	if err := updateList.Validate(); err != nil {
		return err
	}

	return t.repository.UpdateList(userId, listId, updateList)
}
