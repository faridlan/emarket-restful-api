package service

import (
	"context"
	"database/sql"

	"github.com/faridlan/emarket-restful-api/exception"
	"github.com/faridlan/emarket-restful-api/helper"
	"github.com/faridlan/emarket-restful-api/model/domain"
	"github.com/faridlan/emarket-restful-api/model/web"
	"github.com/faridlan/emarket-restful-api/repository"
	"github.com/go-playground/validator"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {

	err := validator.New().Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRolback(tx)

	product := domain.Product{
		Name:       request.Name,
		Price:      request.Price,
		Quantity:   request.Quantity,
		CategoryId: request.CategoryId,
	}

	product = service.ProductRepository.Save(ctx, tx, product)

	return helper.ToProductResponse(product)

}

func (service ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {

	err := validator.New().Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRolback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product.Id = request.Id
	product.Name = request.Name
	product.Price = request.Price
	product.Quantity = request.Quantity
	product.CategoryId = request.CategoryId

	product = service.ProductRepository.Update(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service ProductServiceImpl) Delete(ctx context.Context, productId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRolback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ProductRepository.Delete(ctx, tx, product)
}

func (service ProductServiceImpl) FindById(ctx context.Context, productId int) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRolback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToProductResponse(product)
}

func (service ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRolback(tx)

	products := service.ProductRepository.FindAll(ctx, tx)

	return helper.ToProductResponses(products)
}
