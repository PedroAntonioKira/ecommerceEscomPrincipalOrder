package queries_category

import (
	//Importaciones de go (vienen incluidas al instalar)
	"encoding/json"
	"strconv"
//	"strings"

	//"fmt"

	//importaciones externas (descargadas)
	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/adapters/secundary/database"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func ListOrderQuery(User string, request events.APIGatewayProxyRequest ) (int, string) {

	var fechaDesde, fechaHasta string
	var orderId int
	var page int 

	if len(request.QueryStringParameters["fechaDesde"]) > 0 {
		fechaDesde = request.QueryStringParameters["fechaDesde"]
	}
	if len(request.QueryStringParameters["fechaHasta"]) > 0 {
		fechaHasta = request.QueryStringParameters["fechaHasta"]
	}
	if len(request.QueryStringParameters["page"]) > 0 {
		page, _ = strconv.Atoi(request.QueryStringParameters["page"])
	}
	if len(request.QueryStringParameters["orderId"]) > 0 {
		orderId, _ = strconv.Atoi(request.QueryStringParameters["orderId"])
	}

	result, err2 := database.ListOrderDatabase(User, fechaDesde, fechaHasta, page, orderId)

	if (err2 != nil){
		return 400, "Ocurrió un error al intentar captruar los registros de órdenes del " + fechaDesde + " al " + fechaHasta + " > " + err2.Error()
	}

	Orders, err3 := json.Marshal(result)

	if err3 != nil {
		return 400, "Ocurrió un error al intentar convertir en JSON el registro de Orden"
	}

	return 200, string(Orders)

	

	//return 200, "string(respJson)" +fechaDesde + fechaHasta + orderId + page

}