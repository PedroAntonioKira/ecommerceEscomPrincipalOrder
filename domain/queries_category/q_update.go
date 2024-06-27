package queries_category

import (
	//Importaciones de go (vienen incluidas al instalar)
	"encoding/json"
//	"strconv"

	"fmt"

	//importaciones externas (descargadas)
//	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/domain/entities"
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/adapters/secundary/database"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func UpdateAddressQuery(body string, User string, pathParams int) (int, string) {
	//creamos la variable de la estructura que almacenará todo lo de la categoria relacionada
	var t entities.Address

	// Decodificamos lo que nos viene en el endpoint (json) para guardarlo en la estructura.
	err := json.Unmarshal([]byte(body), &t)

	//Validamos que nos traigan cada uno de los parametros
	if err != nil{
		return 400 , "Error en los datos recibidos " +  err.Error()
	}

	t.AddId = pathParams

	var encontrado bool

	err,encontrado = database.AddressExists(User, t.AddId)

	if(!encontrado) {
		if(err!= nil){
			return 400, "Error al intentar buscar Address para el usuario " + User + " > " + err.Error()
		}else{
			return 400, "No se encuentra un registro de ID de Usuario asociado a esa ID de Address"
		}
	}

	err = database.UpdateAddressQuery(t)

	if(err!= nil){
		return 400, "Ocurrió un error para intentasr realizar la actualización del Address para el ID de Usuario " + User + " > " + err.Error()
	}

	fmt.Println(encontrado)
	fmt.Println(t.AddId)

	return 200, "Update Address OK"

}
