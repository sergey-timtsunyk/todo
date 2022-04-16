package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sergey-timtsunyk/todo/pkg/data"
)

type TodoItemMysql struct {
	db *sqlx.DB
}

func NewTodoItemMysql(db *sqlx.DB) *TodoItemMysql {
	return &TodoItemMysql{
		db: db,
	}
}

func (r *TodoItemMysql) Create(listId uint, input data.Item) (uint, error) {
	query := fmt.Sprintf("INSERT INTO %s (todo_lists_id, item) VALUES (?, ?)", todoItemsTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(listId, input.Item)
	if err != nil {
		return 0, err
	}

	itemId, err := res.LastInsertId()

	return uint(itemId), err
}

func (r *TodoItemMysql) Update(listId uint, itemId uint, input data.UpdateItemInput) error {
	if input.Item == nil {
		return nil
	}
	query := fmt.Sprintf("UPDATE %s SET item = ? WHERE todo_lists_id = ? AND id = ?", todoItemsTable)
	_, err := r.db.Exec(query, input.Item, listId, itemId)

	return err
}

func (r *TodoItemMysql) DoneItem(listId uint, itemId uint) error {
	query := fmt.Sprintf("UPDATE %s SET done = 1, done_at = NOW()  WHERE todo_lists_id = ? AND id = ?", todoItemsTable)
	_, err := r.db.Exec(query, listId, itemId)

	return err
}

func (r *TodoItemMysql) DeleteByIdAndUserId(listId uint, itemId uint) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE todo_lists_id = ? AND id = ?", todoItemsTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(listId, itemId)

	return err
}

func (r *TodoItemMysql) GetAllByListId(listId uint) ([]data.Item, error) {
	var todoItems []data.Item
	query := fmt.Sprintf("SELECT * FROM %s WHERE todo_lists_id = ?", todoItemsTable)
	err := r.db.Select(&todoItems, query, listId)

	return todoItems, err
}

func (r *TodoItemMysql) GetByIdAndListId(listId uint, itemId uint) (data.Item, error) {
	var todoItem data.Item
	query := fmt.Sprintf("SELECT * FROM %s WHERE todo_lists_id = ? AND id = ?", todoItemsTable)
	err := r.db.Get(&todoItem, query, listId, itemId)

	return todoItem, err
}
