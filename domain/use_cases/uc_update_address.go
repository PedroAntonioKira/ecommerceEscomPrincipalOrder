package use_cases

import (
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/domain/repositories"
)

func UpdateAddressUC(body string, user string, pathParams int) (int, string) {
	status := 200
	response := "Vacio Update Category Use Case"

	status, response = repositories.UpdateAddressRepositories(body, user, pathParams)

	return status, response
}