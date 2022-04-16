package data

import "errors"

type SingInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i *UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update struct has no value")
	}

	return nil
}
