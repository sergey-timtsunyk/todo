package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sergey-timtsunyk/todo/pkg/data"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(user data.User) (uint, error)
	GetUserByLoginAndPass(login string, pass string) (data.User, error)
	UpdateLoginDate(user data.User) error
}

type TodoList interface {
	Create(userId uint, list data.List) (uint, error)
	GetAll(userId uint) ([]data.List, error)
	GetByIdAndUserId(userId uint, listId uint) (data.List, error)
	DeleteByIdAndUserId(userId uint, listId uint) error
	UpdateList(userId uint, listId uint, updateList data.UpdateListInput) error
}

type TodoItem interface {
	Create(listId uint, input data.Item) (uint, error)
	Update(listId uint, itemId uint, input data.UpdateItemInput) error
	DoneItem(listId uint, itemId uint) error
	DeleteByIdAndUserId(listId uint, itemId uint) error
	GetAllByListId(listId uint) ([]data.Item, error)
	GetByIdAndListId(listId uint, itemId uint) (data.Item, error)
}

type AuthEvent interface {
	Create(userId uint, uriRequest string, eventName string) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
	AuthEvent
}

func NewRepository(db *sqlx.DB, mongoDb *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
		TodoList:      NewTodoListMysql(db),
		TodoItem:      NewTodoItemMysql(db),
		AuthEvent:     NewAuthEventMongo(mongoDb),
	}
}
