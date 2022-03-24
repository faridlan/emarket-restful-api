package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/emarket-restful-api/helper"
	"github.com/faridlan/emarket-restful-api/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {

	SQL := "insert into users(username,email,password) values (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Email, user.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)

	return user
}

func (repository UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {

	SQL := "select id,username,email from users where (username =? or email =?) and password =?"
	rows, err := tx.QueryContext(ctx, SQL, user.Username, user.Email, user.Password)
	helper.PanicIfError(err)

	defer rows.Close()

	user = domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {

	SQL := "select id,username,email from users where id=?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	var user = domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {

	SQL := "update users set username = ?, email = ?, password = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Username, user.Email, user.Password, user.Id)
	helper.PanicIfError(err)

	return user
}
