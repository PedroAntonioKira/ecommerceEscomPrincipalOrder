package queries_category

import (
	//Importaciones de go (vienen incluidas al instalar)
//	"encoding/json"
	"strconv"

	//"fmt"

	//importaciones externas (descargadas)
//	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/entities"
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalAddress/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/adapters/secundary/database"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func DeleteAddressQuery(body string, User string, pathParams int) (int, string) {
	
	// Validamos que nos hayan pasado un Id valido
	if pathParams == 0 {
		return 400, "Debe de especificar el ID de la categoria a borrar."
	}

	//Validamos que el usuario exista y que este asociada al usuario
	err,encontrado := database.AddressExists(User, pathParams)

	if(!encontrado) {
		if(err!= nil){
			return 400, "Error al intentar buscar Address para el usuario " + User + " > " + err.Error()
		}else{
			return 400, "No se encuentra un registro de ID de Usuario asociado a esa ID de Address"
		}
	}

	//Eliminamos la categoria que corresponde al id
	err = database.DeleteAddressQuery(pathParams)

	// Validamos que no haya surgido un error al eliminar la categoria
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el DELETE del Address" + strconv.Itoa(pathParams) + " > " + err.Error()
	}

	return 200, "Delete Address OK"

}