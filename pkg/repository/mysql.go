package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable     = "users"
	todoListsTable = "todo_lists"
	todoItemsTable = "todo_items"
)

type ConfigMysqlDB struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
}

func NewMysqlDB(cfg ConfigMysqlDB) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
