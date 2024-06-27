package queries_category

import (
	//Importaciones de go (vienen incluidas al instalar)
	"encoding/json"
//	"strconv"

	//"fmt"

	//importaciones externas (descargadas)
//	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/adapters/secundary/database"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func ListAddressQuery(body string, User string) (int, string) {

	// Es el path (CategoryPath), solo que se suele llamar así en un ecommerce

	addr, err := database.ListAddressQuery(User)
	if err != nil{
		return 400, "Ocurrio un error al intentar obtener la lista de direcciones del usuario" + User + " > " + err.Error()
	}

	respJson, err := json.Marshal(addr)
	if err != nil{
		return 500, "Error al formastear los datos de las Addresses como JSON"
	}

	return 200, string(respJson)

}