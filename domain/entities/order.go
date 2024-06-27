// Estructuras para category
package entities

//Estructura para direcciones

type Orders	 struct {
	Order_Id 		int  	`json:"OrderId"`
	Order_UserUUID  string  `json:"OrderUserUUID"`
	Order_AddId		int  	`json:"OrderAddId"`
	Order_Date     	string  `json:"OrderDate"`
	Order_Total     float64 `json:"OrderTotal"`
	OrdersDetails	[]OrdersDetails
}

type OrdersDetails struct{
	OD_Id		int			`json:"odId"`
	OD_OrderId	int			`json:"odOrderId"`
	OD_ProdId	int			`json:"odProdId"`
	OD_Quantity	int			`json:"odQuantity"`
	OD_Price	float64		`json:"odPrice"`
}



type Address struct {
	AddId 			int  	`json:"addId"`
	AddTitle    	string  `json:"addTitle"`
	AddName			string  `json:"addName"`
	AddAddress     	string  `json:"addAddress"`
	AddCity      	string  `json:"addCity"`
	AddState     	string  `json:"addState"`
	AddPostalCode	string  `json:"addPostalCode"`
	AddPhone      	string  `json:"addPhone"`
}