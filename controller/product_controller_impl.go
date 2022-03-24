package controller

import (
	"net/http"
	"strconv"

	"github.com/faridlan/emarket-restful-api/helper"
	"github.com/faridlan/emarket-restful-api/model/web"
	"github.com/faridlan/emarket-restful-api/service"
	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (contoller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	productCreateRequest := web.ProductCreateRequest{}
	helper.ReadFromRequestBody(request, &productCreateRequest)

	productResponse := contoller.ProductService.Create(request.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToRequestBody(writer, webResponse)
}

func (contoller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productUpdateRequest := web.ProductUpdateRequest{}
	helper.ReadFromRequestBody(request, &productUpdateRequest)

	productId := params.ByName("id")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productUpdateRequest.Id = id

	productResponse := contoller.ProductService.Update(request.Context(), productUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToRequestBody(writer, webResponse)
}

func (contoller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("id")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	contoller.ProductService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToRequestBody(writer, webResponse)

}

func (contoller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("id")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productResponse := contoller.ProductService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToRequestBody(writer, webResponse)
}

func (contoller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productResponse := contoller.ProductService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToRequestBody(writer, webResponse)
}
