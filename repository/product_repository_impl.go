package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/emarket-restful-api/helper"
	"github.com/faridlan/emarket-restful-api/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {

	SQL := "insert into products(name_product, price, quantity, category_id) values (?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Quantity, product.CategoryId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)

	return product

}

func (repository ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {

	SQL := "update products set name_product = ?, price = ?, quantity = ?, category_id = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Quantity, product.CategoryId, product.Id)
	helper.PanicIfError(err)

	return product
}

func (repository ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {

	SQL := "delete from products where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)

}

func (repository ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {

	SQL := `select 
					p.id, 
					p.name_product, 
					p.price, 
					p.quantity,
					c.id, 
					c.name_category
					from products as p
					left join categories as c on (p.category_id = c.id) where p.id = ?;`
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	var product = domain.Product{}

	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.CategoryId, &product.CategoryName)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("product not found")
	}

}

func (repository ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {

	SQL := `select 
					p.id, 
					p.name_product,
					p.price,
					p.quantity,
					c.name_category
					from products as p join categories as c on (p.category_id = c.id)`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	products := []domain.Product{}

	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.CategoryName)
		helper.PanicIfError(err)

		products = append(products, product)
	}

	return products
}
