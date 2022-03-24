package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/emarket-restful-api/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
}
