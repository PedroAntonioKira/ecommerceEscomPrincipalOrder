package repositories

import (
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/domain/queries_category"
)

func AddRouterRepositories(body string, user string) (int, string) {
	status := 200
	response := "Vacio Add Category Repositories"

	status, response = queries_category.AddRouterQuery(body, user)

	return status, response
}