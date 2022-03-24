package main

import (
	"net/http"

	"github.com/faridlan/emarket-restful-api/app"
	"github.com/faridlan/emarket-restful-api/controller"
	"github.com/faridlan/emarket-restful-api/exception"
	"github.com/faridlan/emarket-restful-api/helper"
	"github.com/faridlan/emarket-restful-api/middleware"
	"github.com/faridlan/emarket-restful-api/repository"
	"github.com/faridlan/emarket-restful-api/service"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	router := httprouter.New()

	//auth
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db)
	userController := controller.NewUserController(userService)

	router.POST("/api/login", userController.Login)
	router.POST("/api/register", userController.Create)
	router.GET("/api/user/:userId", userController.FindById)
	router.PUT("/api/user/:userId", userController.Update)

	//product
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	router.POST("/api/product", productController.Create)
	router.GET("/api/product/", productController.FindAll)
	router.GET("/api/product/:id", productController.FindById)
	router.PUT("/api/product/:id", productController.Update)
	router.DELETE("/api/product/:id", productController.Delete)

	//category
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router.POST("/api/categories", categoryController.Create)
	router.GET("/api/categories/", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ExceptionError

	// var router1 *httprouter.Router

	// router1 = app.NewRouterUser(userController)
	// router1 = app.NewRouterProduct(productController)
	// router1 = app.NewRouterCategory(categoryController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
