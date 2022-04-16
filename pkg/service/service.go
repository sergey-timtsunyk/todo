package service

import (
	"github.com/sergey-timtsunyk/todo/pkg/data"
	"github.com/sergey-timtsunyk/todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user data.User) (uint, error)
	GenerateToken(login string, password string) (string, error)
	ParserToken(accessToken string) (uint, error)
}

type TodoList interface {
	Create(userId uint, list data.List) (uint, error)
	GetAll(userId uint) ([]data.List, error)
	GetByIdAndUserId(userId uint, listId uint) (data.List, error)
	DeleteByIdAndUserId(userId uint, listId uint) error
	UpdateList(userId uint, listId uint, updateList data.UpdateListInput) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthServer(repository.Authorization),
		TodoList:      NewTodoListService(repository.TodoList),
	}
}
