package service

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/sergey-timtsunyk/todo/pkg/data"
	"github.com/sergey-timtsunyk/todo/pkg/repository"
	"time"
)

const (
	salt       = "jbvkshHVkjbljhv"
	signingKey = "kjvbkgc&*gJKHVIytr"
	tokenTtl   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId uint64 `json:"user_id"`
}

type AuthService struct {
	repository repository.Authorization
}

func NewAuthServer(repository repository.Authorization) *AuthService {
	return &AuthService{repository: repository}
}

func (as *AuthService) CreateUser(user data.User) (uint64, error) {
	user.Password = generatePasswordHash(user.Password)

	return as.repository.CreateUser(user)
}

func (as *AuthService) GenerateToken(login string, password string) (string, error) {
	user, err := as.repository.GetUserByLoginAndPass(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTtl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	if err := as.repository.UpdateLoginDate(user); err != nil {
		return "", err
	}

	return token.SignedString([]byte(signingKey))
}

func (as *AuthService) ParserToken(accessToken string) (uint64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
