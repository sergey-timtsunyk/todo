package data

import "time"

type List struct {
	Id          uint64    `json:"id" db:"id"`
	UserId      uint64    `json:"user_id" db:"user_id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"description" db:"description" binding:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Item struct {
	Id          uint64    `json:"id"`
	TodoListsId uint64    `json:"todo_lists_id"`
	Item        string    `json:"item"`
	Done        bool      `json:"done"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
	DoneAt      time.Time `json:"done_at"`
}
