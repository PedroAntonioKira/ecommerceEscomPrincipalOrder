package repositories

import (
	//Importaciones de go (vienen incluidas al instalar)
	"fmt"
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/domain/queries_category"
)

func DeleteAddressRepositories(body string, user string, pathParams int) (int, string) {
	status := 200
	response := "Vacio Delete Category Repositories"

	fmt.Println("Entramos a DeleteCategoryRepositories")
	status, response = queries_category.DeleteAddressQuery(body, user, pathParams)

	return status, response
}