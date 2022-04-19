package service

import (
	"fmt"
	"github.com/sergey-timtsunyk/todo/pkg/data"
	"github.com/sergey-timtsunyk/todo/pkg/repository"
)

type AuthEventService struct {
	repository repository.AuthEvent
}

func NewAuthEventService(repository repository.AuthEvent) *AuthEventService {
	return &AuthEventService{repository: repository}
}

func (s *AuthEventService) AddVerificationEvent(userId uint, method string, uri string) error {
	uriRequest := fmt.Sprintf("%s: %s", method, uri)
	return s.repository.Create(userId, uriRequest, data.VerificationEvent)
}

func (s *AuthEventService) AddAuthenticationEvent(userId uint, method string, uri string) error {
	uriRequest := fmt.Sprintf("%s: %s", method, uri)
	return s.repository.Create(userId, uriRequest, data.AuthenticationEvent)
}

func (s *AuthEventService) AddAuthorizationEvent(userId uint, method string, uri string) error {
	uriRequest := fmt.Sprintf("%s: %s", method, uri)
	return s.repository.Create(userId, uriRequest, data.AuthorizationEvent)
}
