package service

import (
	todo "github.com/zemavong/TodoList"
	"github.com/zemavong/TodoList/pkg/repository"
)

type todoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *todoListService {
	return &todoListService{repo: repo}
}

func (s *todoListService) Create(userId int, list todo.Todo) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *todoListService) GetAll(userId int) ([]todo.Todo, error) {
	return s.repo.GetAll(userId)
}

func (s *todoListService) GetById(userId, listId int) (todo.Todo, error) {
	return s.repo.GetById(userId, listId)
}

func (s *todoListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *todoListService) Update(userId, listId int, input todo.UpdateListInput) error {
	return s.repo.Update(userId, listId, input)
}
