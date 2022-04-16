package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sergey-timtsunyk/todo/pkg/data"
	"time"
)

type AuthMysql struct {
	db *sqlx.DB
}

func NewAuthMysql(db *sqlx.DB) *AuthMysql {
	return &AuthMysql{db: db}
}

func (r *AuthMysql) CreateUser(user data.User) (uint64, error) {
	query := fmt.Sprintf("INSERT INTO %s (full_name, login, password) VALUES (?, ?, ?)", usersTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(user.Name, user.Login, user.Password)
	if err != nil {
		return 0, err
	}
	userId, err := res.LastInsertId()

	return uint64(userId), err
}

func (r *AuthMysql) GetUserByLoginAndPass(login string, pass string) (data.User, error) {
	var user data.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE login = ? AND password = ?", usersTable)
	err := r.db.Get(&user, query, login, pass)

	return user, err
}

func (r *AuthMysql) UpdateLoginDate(user data.User) error {
	query := fmt.Sprintf("UPDATE %s SET login_at = ? WHERE id = ?", usersTable)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(time.Now(), user.Id); err != nil {
		return err
	}

	return nil
}
