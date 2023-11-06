package controllers

import (
	"fmt"
	"strconv"

	auth "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	servs "github.com/sabrs0/bmstu-web/src/internal/business/services"
	"github.com/sabrs0/bmstu-web/src/internal/my_errors"
)

type TokenGenerator interface {
	GenerateToken(role, id string) (string, error)
}

type LoginController struct {
	FS *servs.FoundationService
	US *servs.UserService
	AS TokenGenerator //для него какие параметры при создании (например тип аутентификации, но пока пусто)
}

func NewLoginController(FR servs.IFoundationRepository,
	UR servs.IUserRepository,
	generator TokenGenerator) *LoginController {
	return &LoginController{
		FS: servs.NewFoundationService(FR),
		US: servs.NewUserService(UR),
		AS: generator,
	}
}
func (ctrl *LoginController) Login(params auth.Params) (string, error) {
	var token string /*@TODO: remove soon ofk*/

	switch params.Role {
	case "User":
		u, err := ctrl.US.GetByLogin(params.Login)
		if err != nil {
			return "", fmt.Errorf("%w: %w", my_errors.ErrorAuth, err)
		}
		if u.Password != params.Password {
			return "", fmt.Errorf("%w: неверный пароль", my_errors.ErrorAuth)
		}
		//set token
		token, err = ctrl.AS.GenerateToken(params.Role, strconv.FormatUint(u.Id, 10))
		if err != nil {
			return "", fmt.Errorf("%w: %w", my_errors.ErrorAuth, err)
		}
	case "Foundation":
		f, err := ctrl.FS.GetByLogin(params.Login)
		if err != nil {
			return "", fmt.Errorf("%w: %w", my_errors.ErrorAuth, err)
		}
		if f.Password != params.Password {
			return "", fmt.Errorf("%w: неверный пароль", my_errors.ErrorAuth)
		}
		token, err = ctrl.AS.GenerateToken(params.Role, strconv.FormatUint(f.Id, 10))
		if err != nil {
			return "", fmt.Errorf("%w: %w", my_errors.ErrorAuth, err)
		}
	case "Admin":
		if params.Login != "ADMIN" && params.Password != "ADMIN" {
			return "", fmt.Errorf("%w: неверный пароль", my_errors.ErrorAuth)
		}
		//set token
		var err error
		token, err = ctrl.AS.GenerateToken(params.Role, "0")
		if err != nil {
			return "", fmt.Errorf("%w: %w", my_errors.ErrorAuth, err)
		}
	}
	return token, nil
}
