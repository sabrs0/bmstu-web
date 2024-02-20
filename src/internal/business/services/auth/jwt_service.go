package auth

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var jwtKey = []byte("supersecretkey")

type JWTClaims struct {
	EntityId string `json:"ent_id"`
	Role     string `json:"role"`
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
		EntityId: id,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func (service *JWTService) ValidateToken(signedToken string) (*jwt.Token, error) {
	//claims := JWTClaims{}
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		err = fmt.Errorf("TOKEN IS %s: %w", signedToken, err)
		return nil, err
	}
	if !token.Valid {
		err = fmt.Errorf("TOKEN IS %s: %w", signedToken, err)
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New("Не удалось распарсить claims")
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("Токен просрочен")
	}
	//логика валидации токена на разрешение операции
	return token, nil

}
func (service *JWTService) ValidateRequest(r *http.Request, roleToAccept []string) error {
	strToken := r.Header.Get("Authorization") //strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ") //
	token, err := service.ValidateToken(strToken)
	if err != nil {
		return fmt.Errorf("Невалидный токен %s. %w", strToken, err)
	}
	claims, _ := token.Claims.(*JWTClaims)
	flag := 0
	for _, role := range roleToAccept {
		if claims.Role == role {
			flag = 1
			break
		}
	}
	if flag == 0 {
		return fmt.Errorf("Неверная роль пользователя %s. Должна быть %v", claims.Role, roleToAccept)
	}
	if claims.Role == "Admin" {
		return nil
	}
	vars := mux.Vars(r)
	if vars == nil || vars["id"] == "" {
		return nil
	}
	/*id := vars["id"]
	if claims.EntityId != id {
		return fmt.Errorf("Неверный  id пользователя")
	}*/
	return nil
}
