package app

import (
	"github.com/faridlan/emarket-restful-api/controller"
	"github.com/faridlan/emarket-restful-api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouterCategory(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	//category
	router.POST("/api/categories", categoryController.Create)
	router.GET("/api/categories/", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ExceptionError

	return router
}
func NewRouterProduct(productController controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	//product
	router.POST("/api/product", productController.Create)
	router.GET("/api/product/", productController.FindAll)
	router.GET("/api/product/:id", productController.FindById)
	router.PUT("/api/product/:id", productController.Update)
	router.DELETE("/api/product/:id", productController.Delete)

	router.PanicHandler = exception.ExceptionError

	return router
}
func NewRouterUser(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	//auth
	router.POST("/api/login", userController.Login)
	router.POST("/api/register", userController.Create)
	router.GET("/api/user/:userId", userController.FindById)
	router.PUT("/api/user/:userId", userController.Update)

	router.PanicHandler = exception.ExceptionError

	return router
}
