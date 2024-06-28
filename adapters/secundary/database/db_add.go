package database

import (
	//Importaciones de go (vienen incluidas al instalar)
	"database/sql"
	"fmt"
	"strconv"
//	"strings"

	//"strings"

	//importaciones externas (descargadas)
	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/domain/entities"
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalAddress/utils"
)

func AddRouterDatabase(o entities.Orders)  (int64, error) {
	fmt.Println("Comienza Registro de Order")

	//Nos conectamos a la base de datos
	err := secundary.DbConnect()
	fmt.Println("Despues de la conexión")
	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return 0, err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer secundary.Db.Close()

	//Declaramos la sentencia SQL para insertar el Producto.
	sentencia := "INSERT INTO orders (Order_UserUUID, Order_Total, Order_AddId) VALUES ('"
	sentencia += o.Order_UserUUID + "', " + strconv.FormatFloat(o.Order_Total, 'f', -1, 64) + ", " + strconv.Itoa(o.Order_AddId) + " )"

	var result sql.Result

	fmt.Println(sentencia)

	//Ejecutamos la sentencia SQL
	result, err = secundary.Db.Exec(sentencia)

	//Verificamos no haya existido un error al ejecutar la sentencia SQL
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()

	if err2 != nil {
		fmt.Println(err2.Error())
		return 0, err
	}

	for _, od := range o.OrdersDetails {
		sentencia = "INSERT INTO orders_detail (OD_OrderId, OD_ProdId, Od_Quantity, OD_Price) VALUES (" + strconv.Itoa(int(LastInsertId))
		sentencia += ", " + strconv.Itoa(od.OD_ProdId) + ", " + strconv.Itoa(od.OD_Quantity) + ", " + strconv.FormatFloat(od.OD_Price, 'f', -1, 64) + " )"
		fmt.Println(sentencia)

		//Ejecutamos la sentencia SQL
		_, err = secundary.Db.Exec(sentencia)

		//Verificamos no haya existido un error al ejecutar la sentencia SQL
		if err != nil {
			fmt.Println(err.Error())
			return 0, err
		}
	}

	fmt.Println("Insert Routers > Ejecución Exitosa")

	return LastInsertId,nil

}



