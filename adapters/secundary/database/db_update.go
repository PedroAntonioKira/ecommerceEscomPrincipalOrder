package database

import (
	//Importaciones de go (vienen incluidas al instalar)
//	"database/sql"
	"fmt"
	"strconv"
	"strings"

	//"strings"

	//importaciones externas (descargadas)
	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalOrder/domain/entities"
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/utils"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
)

func UpdateAddressQuery(addr entities.Address) error {

	fmt.Println("Comienza Registro de UpdateAddressQuery")

	//Nos conectamos a la base de datos
	err := secundary.DbConnect()
	fmt.Println("Despues de la conexión")
	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer secundary.Db.Close()

	sentencia := "UPDATE addresses SET "

	if addr.AddAddress != ""{
		sentencia += "Add_Address = '" + addr.AddAddress + "', "
	}

	if addr.AddCity != ""{
		sentencia += "Add_City = '" + addr.AddCity + "', "
	}

	if addr.AddName != ""{
		sentencia += "Add_Name = '" + addr.AddName + "', "
	}

	if addr.AddPhone != ""{
		sentencia += "Add_Phone = '" + addr.AddPhone + "', "
	}

	if addr.AddPostalCode != ""{
		sentencia += "Add_PostalCode = '" + addr.AddPostalCode + "', "
	}

	if addr.AddState != ""{
		sentencia += "Add_State = '" + addr.AddState + "', "
	}

	if addr.AddTitle != ""{
		sentencia += "Add_Title = '" + addr.AddTitle + "', "
	}

	sentencia, _ = strings.CutSuffix(sentencia, ", ")

	sentencia += " WHERE Add_Id = " + strconv.Itoa(addr.AddId)

	fmt.Println(sentencia)

	_, err = secundary.Db.Query(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Address > Ejecución Exitosa")

	return nil

}

func AddressExists(User string, pathParams int) (error, bool){
	fmt.Println("Comienza AddressExists")

	//Nos conectamos a la base de datos
	err := secundary.DbConnect()
	fmt.Println("Despues de la conexión")
	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return err, false
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer secundary.Db.Close()

	sentencia := "SELECT 1 FROM addresses WHERE Add_Id = " + strconv.Itoa(pathParams) + " AND  Add_UserId = '" + User + "'"

	fmt.Println(sentencia)

	rows, err02 := secundary.Db.Query(sentencia)

	if err02 != nil {
		return err02, false
	}

	var valor string
	rows.Next()
	rows.Scan(&valor)

	fmt.Println("AddressExists > Ejecución Exitosa - valor devuelto_ " + valor)
	
	if( valor == "1"){
		return nil, true
	}else{
		return nil, false
	}

}