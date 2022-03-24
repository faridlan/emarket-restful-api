package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/faridlan/emarket-restful-api/helper"
	"github.com/faridlan/emarket-restful-api/model/domain"
	"github.com/faridlan/emarket-restful-api/model/web"
	"github.com/golang-jwt/jwt/v4"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	authorizationHeader := request.Header.Get("Authorization")

	if request.URL.Path == "/api/login" || request.URL.Path == "/api/register" {
		middleware.Handler.ServeHTTP(writer, request)
		return
	}

	if !strings.Contains(authorizationHeader, "Bearer") {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToRequestBody(writer, webResponse)
		return
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	var claim = &domain.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != domain.JwtSigningMEethod {
			return nil, fmt.Errorf("signing method invalid")
		}
		return domain.JwtSecret, nil
	})

	if err != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}

		helper.WriteToRequestBody(writer, webResponse)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
		}

		helper.WriteToRequestBody(writer, webResponse)
		writer.WriteHeader(http.StatusBadRequest)
		return
	} else {
		middleware.Handler.ServeHTTP(writer, request)
	}

}
