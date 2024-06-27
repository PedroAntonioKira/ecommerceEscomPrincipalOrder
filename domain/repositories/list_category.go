package repositories

import (
	//Importaciones de go (vienen incluidas al instalar)
//	"encoding/json"
//	"strconv"

	"fmt"

	//importaciones externas (descargadas)
//	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/domain/queries_category"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func ListAddressRepositories(body string, User string) (int, string) {

	status := 200
	response := "Vacio"


	fmt.Println("Entramos a ListCategoryRepositories")
	status, response = queries_category.ListAddressQuery(body, User)


	return status, response
}