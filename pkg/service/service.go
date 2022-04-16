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
	Create(userId uint, listId uint, inputItem data.Item) (uint, error)
	Update(userId uint, listId uint, itemId uint, updateItem data.UpdateItemInput) error
	DoneItem(userId uint, listId uint, itemId uint) error
	DeleteByIdAndUserId(userId uint, listId uint, itemId uint) error
	GetAllByListId(userId uint, listId uint) ([]data.Item, error)
	GetByIdAndListId(userId uint, listId uint, itemId uint) (data.Item, error)
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
		TodoItem:      NewTodoItemService(repository.TodoItem, repository.TodoList),
	}
}
