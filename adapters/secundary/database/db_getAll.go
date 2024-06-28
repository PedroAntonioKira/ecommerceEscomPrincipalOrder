package database

import (
	//Importaciones de go (vienen incluidas al instalar)
	"database/sql"
	"fmt"
	"strconv"
//	"strings"

	//importaciones externas (descargadas)
	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/domain/entities"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/t|ools"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
)

func ListOrderDatabase(User string, fechaDesde string, fechaHasta string, page int, orderId int) ([]entities.Orders, error) {
	fmt.Println("Comienza ListOrdersQuery")

	//Creamos la variable que almacenara cada registro devuelto de cada categoria de la base de datos.
	var Orders []entities.Orders

	//Creamos el inciio de la sentencia SQL
	sentencia := "SELECT Order_Id, Order_UserUUID, Order_AddId, Order_Date, Order_Total FROM orders"

	if (orderId > 0) {
		sentencia += " WHERE Order_Id = " + strconv.Itoa(orderId)
	} else {
		offset := 0
		if(page == 0){
			page = 1
		}
		if(page > 1){
			offset = ( 10 * (page - 1))
		}

		if len(fechaHasta) == 10{
			fechaHasta += " 23:59:59"
		}

		var where string
		var whereUser string = "Order_UserUUID = '" + User + "'"

		if len(fechaDesde) > 0 && len(fechaHasta) > 0 {
			where += " WHERE Order_Date BETWEEN '" + fechaDesde + "' AND '" + fechaHasta
		}
		if len(where) > 0 {
			where += " AND " + whereUser
		} else{
			where += " WHERE " + whereUser
		}

		limit := " LIMIT 10 "
		if offset > 0 {
			limit += " OFFSET " + strconv.Itoa(offset)
		}

		sentencia += where + limit
	}

	fmt.Println("Imprimimos la sentencia para ver la info:")
	fmt.Println(sentencia)

	//Nos conectamos a la base de datos
	err := secundary.DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return Orders, err
	}

	//Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer secundary.Db.Close()

	var rows *sql.Rows

	rows, err = secundary.Db.Query(sentencia)
	
	//Verificamos no haya existido un error al ejecutar la sentencia SQL
	if err != nil {
		fmt.Println("Fallo depués de ejecutar los datos sql")
		fmt.Println(err.Error())
		return Orders, err
	}

	defer rows.Close()

	fmt.Println("Aqui pasamos el defere y apenas entraremos al for")

	for rows.Next() {
		fmt.Println("Aqui Entramos al for")
		var Order entities.Orders
		var OrderDate sql.NullString
		var OrderAddId sql.NullInt32

		fmt.Println("Llegamops al for antes de guardar los datos")
		err := rows.Scan(&Order.Order_Id, &Order.Order_UserUUID, &OrderAddId, &OrderDate, &Order.Order_Total)

		fmt.Println("Terminamos de leer los datos")
		if err != nil {
			fmt.Println("Aqui peude estar el error 001")
			fmt.Println(err.Error())
			return Orders, err
		}

		Order.Order_Date = OrderDate.String
		Order.Order_AddId = int(OrderAddId.Int32)

		var rowsD *sql.Rows

		sentenciaD := "SELECT OD_Id, OD_ProdId, OD_Quantity, OD_Price FROM orders_detail WHERE OD_OrderId = " + strconv.Itoa(Order.Order_Id) 

		rowsD, err = secundary.Db.Query(sentenciaD)

		//Verificamos no haya existido un error al ejecutar la sentencia SQL
		if err != nil {
			fmt.Println("Fallo depués de ejecutar los datos sql 02")
			fmt.Println(err.Error())
			return Orders, err
		}

		for rowsD.Next(){
			var OD_Id int64
			var OD_ProdId int64
			var OD_Quantity int64
			var OD_Price float64

			err := rowsD.Scan(&OD_Id, &OD_ProdId, &OD_Quantity, &OD_Price)

			if err != nil {
				fmt.Println("Aqui peude estar el error 002")
				fmt.Println(err.Error())
				return Orders, err
			}

			var od entities.OrdersDetails
			od.OD_Id = int(OD_Id)
			od.OD_ProdId = int(OD_ProdId)
			od.OD_Quantity = int(OD_Quantity)
			od.OD_Price = OD_Price

			Order.OrdersDetails = append(Order.OrdersDetails, od)

		}

		Orders = append(Orders, Order)

		rowsD.Close()
	}
	fmt.Println("Select Orders > Ejecución Exitosa")
	fmt.Println("Cambiar texto")

	return Orders, nil
}