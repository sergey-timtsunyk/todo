package data

type User struct {
	Id       uint64 `json:"id" db:"id"`
	Name     string `json:"fullName" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
