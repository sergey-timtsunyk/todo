package service

import (
	"github.com/sergey-timtsunyk/todo/pkg/data"
	"github.com/sergey-timtsunyk/todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user data.User) (uint64, error)
	GenerateToken(login string, password string) (string, error)
	ParserToken(accessToken string) (uint64, error)
}

type TodoList interface {
	Create(userId uint64, list data.List) (uint64, error)
	GetAll(userId uint64) ([]data.List, error)
	GetByIdAndUserId(userId uint64, listId uint64) (data.List, error)
	DeleteByIdAndUserId(userId uint64, listId uint64) error
	UpdateList(userId uint64, listId uint64, updateList data.UpdateListInput) error
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
