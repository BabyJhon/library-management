package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/BabyJhon/library-managment/internal/entity"
	"github.com/BabyJhon/library-managment/internal/repo"
	"github.com/golang-jwt/jwt/v5"
)

const (
	salt       = "asdklijasd"
	TokenTTL   = 12 * time.Hour
	signingKey = "das123890op123eawd"
)

type AuthService struct {
	repo repo.Authorization
}

type tokenClaims struct {
	jwt.RegisteredClaims
	Id int `json:"id"`
}

func NewAuthService(repo repo.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateAdmin(c context.Context, input entity.Admin) (int, error) {
	input.Password = a.generatePasswordHash(input.Password)

	return a.repo.CreateAdmin(c, input)
}

func (a *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (a *AuthService) GenerateToken(ctx context.Context, userName, password string) (string, error) {
	admin, err := a.repo.GetAdmin(ctx, userName, a.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Id: admin.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (a *AuthService) Parsetoken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, nil
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.Id, nil
}
