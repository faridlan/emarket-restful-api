package helper

import (
	"github.com/faridlan/emarket-restful-api/model/domain"
	"github.com/faridlan/emarket-restful-api/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}

func ResponseToken(claims domain.Claims) domain.Claims {
	return domain.Claims{
		Id:               claims.Id,
		Username:         claims.Username,
		Email:            claims.Email,
		RegisteredClaims: claims.RegisteredClaims,
	}
}

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:           product.Id,
		Name:         product.Name,
		Price:        product.Price,
		Quantity:     product.Quantity,
		CategoryId:   product.CategoryId,
		CategoryName: product.CategoryName,
	}
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}

	return productResponses
}

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {

	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}
