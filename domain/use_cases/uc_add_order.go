package use_cases

import (
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/domain/repositories"
)

func AddRouterUC(body string, user string) (int, string) {
	status := 200
	response := "Vacio Add Category Use Case"

	status, response = repositories.AddRouterRepositories(body, user)

	return status, response
}