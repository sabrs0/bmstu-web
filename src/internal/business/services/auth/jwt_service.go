package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("supersecretkey")

type JWTClaims struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type JWTService struct {
}

func NewJWTService() *JWTService {
	return &JWTService{}
}
func (service *JWTService) GenerateToken(role, id string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 1)
	claims := &JWTClaims{
		Id:   id,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(signedToken string) error {
	//claims := JWTClaims{}
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return errors.New("Не удалось распарсить claims")
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return errors.New("Токен просрочен")
	}
	//логика валидации токена на разрешение операции
	return nil

}
