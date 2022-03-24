package service

import (
	"context"

	"github.com/faridlan/emarket-restful-api/model/domain"
	"github.com/faridlan/emarket-restful-api/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Login(ctx context.Context, request web.UserCreateRequest) domain.Claims
	FindById(ctx context.Context, userId int) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
}
