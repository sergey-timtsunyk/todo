package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sergey-timtsunyk/todo/pkg/data"
	"strings"
)

type TodoListMysql struct {
	db *sqlx.DB
}

func NewTodoListMysql(db *sqlx.DB) *TodoListMysql {
	return &TodoListMysql{db: db}
}

func (r *TodoListMysql) Create(userId uint64, list data.List) (uint64, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, title, description) VALUES (?, ?, ?)", todoListsTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(userId, list.Title, list.Description)
	if err != nil {
		return 0, err
	}

	listId, err := res.LastInsertId()

	return uint64(listId), err
}

func (r *TodoListMysql) GetAll(userId uint64) ([]data.List, error) {
	var todoLists []data.List
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = ?", todoListsTable)
	err := r.db.Select(&todoLists, query, userId)

	return todoLists, err
}

func (r *TodoListMysql) GetByIdAndUserId(userId uint64, listId uint64) (data.List, error) {
	var todoList data.List
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = ? AND id = ?", todoListsTable)
	err := r.db.Get(&todoList, query, userId, listId)

	return todoList, err
}

func (r *TodoListMysql) DeleteByIdAndUserId(userId uint64, listId uint64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = ? AND id = ?", todoListsTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userId, listId)

	return err
}

func (r *TodoListMysql) UpdateList(userId uint64, listId uint64, updateList data.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	if updateList.Title != nil {
		setValues = append(setValues, "title=?")
		args = append(args, &updateList.Title)
	}
	if updateList.Description != nil {
		setValues = append(setValues, "description=?")
		args = append(args, &updateList.Description)
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE user_id = ? AND id = ?", todoListsTable, setQuery)
	args = append(args, userId, listId)
	_, err := r.db.Exec(query, args...)

	return err
}
