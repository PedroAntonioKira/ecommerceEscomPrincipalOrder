package use_cases

import (
	//Importaciones de go (vienen incluidas al instalar)
	"fmt"
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/domain/repositories"
)

func DeleteAddressUC(body string, user string, pathParams int) (int, string) {

	fmt.Println("Entramos a DeleteCategoryUC")
	status := 200
	response := "Vacio Delete Category Use Case"

	status, response = repositories.DeleteAddressRepositories(body, user, pathParams)

	return status, response
}