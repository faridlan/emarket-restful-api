package service

import (
	"context"
	"database/sql"

	"github.com/faridlan/emarket-restful-api/helper"
	"github.com/faridlan/emarket-restful-api/model/domain"
	"github.com/faridlan/emarket-restful-api/model/web"
	"github.com/faridlan/emarket-restful-api/repository"
	"github.com/golang-jwt/jwt/v4"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
	}
}

func (service UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRolback(tx)

	user := domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	user = service.UserRepository.Save(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service UserServiceImpl) Login(ctx context.Context, request web.UserCreateRequest) domain.Claims {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRolback(tx)

	user := domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	user, err = service.UserRepository.Login(ctx, tx, user)
	helper.PanicIfError(err)

	claim := domain.Claims{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(domain.ExpiredTime),
		},
	}

	return helper.ResponseToken(claim)
}

func (service UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRolback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}

func (service UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRolback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	user.Username = request.Username
	user.Email = request.Email
	user.Password = request.Password

	user = service.UserRepository.Update(ctx, tx, user)

	return helper.ToUserResponse(user)

}
