package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sergey-timtsunyk/todo/pkg/data"
)

type Authorization interface {
	CreateUser(user data.User) (uint64, error)
	GetUserByLoginAndPass(login string, pass string) (data.User, error)
	UpdateLoginDate(user data.User) error
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

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
		TodoList:      NewTodoListMysql(db),
	}
}
