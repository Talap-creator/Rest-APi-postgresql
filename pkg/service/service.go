package service

import "Rest-APi-postgresql/pkg/repository"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service interface {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
