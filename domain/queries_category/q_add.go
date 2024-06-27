package queries_category

import (
	//Importaciones de go (vienen incluidas al instalar)
	"encoding/json"
	"strconv"
//	"fmt"

	//importaciones externas (descargadas)
	//"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/domain/entities"
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalAddress/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/adapters/secundary/database"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func AddRouterQuery(body string, User string) (int, string) {
	
	var o entities.Orders
	err := json.Unmarshal([]byte(body), &o)

	if err != nil{
		return 400, "Error en los datos recibidos " + err.Error()
	}

	o.Order_UserUUID = User

	Ok, menssage := ValidOrder(o)

	if !Ok{
		return 400, menssage
	}

	result, err2 := database.AddRouterDatabase(o)

	if err2 != nil{
		return 400, "Ocurrio un Error al intentar realizar el registro de la orden " + err2.Error()
	}

	return 200, "{ OrderID: " + strconv.Itoa(int(result)) + "}"
}

func ValidOrder (o entities.Orders) (bool, string){
	if o.Order_Total == 0{
		return false, "Debe indicar el total de la orden"
	}

	count := 0
	for _, od := range o.OrdersDetails {
		if(od.OD_ProdId == 0){
			return false, "Debe indicar el ID del producto en el detalle de la orden"
		}
		if(od.OD_Quantity == 0){
			return false, "Debe indicar la cantidad del producto en el detalle de la orden"
		}
		count++
	}

	return true, ""
}