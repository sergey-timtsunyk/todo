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

type UpdateItemInput struct {
	Item *string ` json:"item"`
	Done *bool   `json:"done"`
}

func (i *UpdateItemInput) Validate() error {
	if i.Item == nil && i.Done == nil {
		return errors.New("update struct has no value")
	}

	return nil
}
