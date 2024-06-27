package repositories

import (
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/domain/queries_category"
)

func UpdateAddressRepositories(body string, user string, pathParams int) (int, string) {
	status := 200
	response := "Vacio Update Category Repositories"

	status, response = queries_category.UpdateAddressQuery(body, user, pathParams)

	return status, response
}