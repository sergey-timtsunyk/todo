package service

import (
	"github.com/sergey-timtsunyk/todo/pkg/data"
	"github.com/sergey-timtsunyk/todo/pkg/repository"
)

type TodoItemService struct {
	repository     repository.TodoItem
	listRepository repository.TodoList
}

func NewTodoItemService(repository repository.TodoItem, listRepository repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repository:     repository,
		listRepository: listRepository,
	}
}

func (s *TodoItemService) Create(userId uint, listId uint, inputItem data.Item) (uint, error) {
	if _, err := s.listRepository.GetByIdAndUserId(userId, listId); err != nil {
		return 0, err
	}

	return s.repository.Create(listId, inputItem)
}

func (s *TodoItemService) Update(userId uint, listId uint, itemId uint, updateItem data.UpdateItemInput) error {
	if _, err := s.listRepository.GetByIdAndUserId(userId, listId); err != nil {
		return err
	}

	if err := updateItem.Validate(); err != nil {
		return err
	}

	if updateItem.Done != nil && *updateItem.Done {
		if err := s.DoneItem(userId, listId, itemId); err != nil {
			return err
		}
	}

	return s.repository.Update(listId, itemId, updateItem)
}

func (s *TodoItemService) DoneItem(userId uint, listId uint, itemId uint) error {
	if _, err := s.listRepository.GetByIdAndUserId(userId, listId); err != nil {
		return err
	}

	return s.repository.DoneItem(listId, itemId)
}

func (s *TodoItemService) DeleteByIdAndUserId(userId uint, listId uint, itemId uint) error {
	if _, err := s.listRepository.GetByIdAndUserId(userId, listId); err != nil {
		return err
	}

	return s.repository.DeleteByIdAndUserId(listId, itemId)
}

func (s *TodoItemService) GetAllByListId(userId uint, listId uint) ([]data.Item, error) {
	if _, err := s.listRepository.GetByIdAndUserId(userId, listId); err != nil {
		return nil, err
	}

	return s.repository.GetAllByListId(listId)
}

func (s *TodoItemService) GetByIdAndListId(userId uint, listId uint, itemId uint) (data.Item, error) {
	return s.repository.GetByIdAndListId(listId, itemId)
}
