package tax

import (
	"log"

	cRouter "github.com/bagusandrian/mini-api/src/common/router"
	"github.com/bagusandrian/mini-api/src/config"
	"github.com/bagusandrian/mini-api/src/db"
	"github.com/jmoiron/sqlx"
	jsoniter "github.com/json-iterator/go"
)

var jsoni = jsoniter.ConfigCompatibleWithStandardLibrary

type Module struct {
	Conf         *config.Config
	queries      *Queries
	CoreMasterDB *sqlx.DB
	CoreSlaveDB  *sqlx.DB
}

func NewModule(c *config.Config) *Module {

	m := &Module{
		Conf: c,
	}

	m.CoreMasterDB = db.Get("CoreMaster")
	if m.CoreMasterDB == nil {
		log.Println("[ERROR] connecting to CoreMaster")
		return nil
	}

	m.CoreSlaveDB = db.Get("CoreSlave")
	if m.CoreSlaveDB == nil {
		log.Panic("[ERROR] connecting to CoreSlave")
		return nil
	}

	m.queries = NewQueries(m.CoreMasterDB, m.CoreSlaveDB)

	return m
}

func RegisterRoutes(r *cRouter.MyRouter, mdle *Module) {

	// Get information list of tax
	// endpoint: localhost:9090/mini-api/tax/list
	// HTTP Method: GET
	// example response:
	// {
	// 	"header": {
	// 		"process_time": 2.135996,
	// 		"messages": null,
	// 		"reason": "",
	// 		"error_code": 0
	// 	},
	// 	"data": {
	// 		"List": [
	// 			{
	// 				"IDTax": 1,
	// 				"Type": "food"
	// 			},
	// 			{
	// 				"IDTax": 2,
	// 				"Type": "tobacco"
	// 			},
	// 			{
	// 				"IDTax": 3,
	// 				"Type": "entertainment"
	// 			}
	// 		]
	// 	},
	// 	"Message": "Success"
	// }
	r.GET("/tax/list", mdle.ListTaxHandler)

	// register user
	// endpoint: localhost:9090/mini-api/user/register
	// HTTP Method: POST
	//	+-------------------------------------------------------------------+
	// 	| Params	| type		| null	| Default	| Description			|
	//	--------------------------------------------------------------------+
	// 	| name		| string	| NO	| 			| name value for user	|
	//	--------------------------------------------------------------------+
	// example response:
	// {
	// 	"header": {
	// 		"process_time": 8.36921,
	// 		"messages": null,
	// 		"reason": "",
	// 		"error_code": 0
	// 	},
	// 	"data": {
	// 		"ID": 5,
	// 		"Name": "test post"
	// 	},
	// 	"Message": "Success"
	// }
	r.POST("/user/register", mdle.RegisterUserHandler)

	// Get list of users
	// endpoint: localhost:9090/mini-api/user/list
	// HTTP Method: GET
	// example response:
	// {
	// 	"header": {
	// 		"process_time": 1.610863,
	// 		"messages": null,
	// 		"reason": "",
	// 		"error_code": 0
	// 	},
	// 	"data": [
	// 		{
	// 			"ID": 1,
	// 			"Name": "user_test"
	// 		},
	// 		{
	// 			"ID": 3,
	// 			"Name": "test post"
	// 		},
	// 		{
	// 			"ID": 4,
	// 			"Name": "test post"
	// 		},
	// 		{
	// 			"ID": 5,
	// 			"Name": "test post"
	// 		}
	// 	],
	// 	"Message": "Success"
	// }
	r.GET("/user/list", mdle.GetListUserHandler)

	// Register new product
	// endpoint: localhost:9090/mini-api/product/register
	// HTTP Method: POST
	//	+-----------------------------------------------------------------------+
	// 	| Params		| type		| null	| Default	| Description			|
	//	------------------------------------------------------------------------+
	// 	| product_name	| string	| NO	| 			| Product name			|
	//	------------------------------------------------------------------------+
	// 	| price			| float64	| NO	| 			| Product price			|
	//	------------------------------------------------------------------------+
	// 	| tax_code_id	| int		| NO	| 			| Code of tax. U can get|
	// 	|				|			|		|			| from tax/list			|
	//	------------------------------------------------------------------------+
	// example response:
	// {
	// 	"header": {
	// 		"process_time": 2.399198,
	// 		"messages": null,
	// 		"reason": "",
	// 		"error_code": 0
	// 	},
	// 	"data": {
	// 		"ID": 10,
	// 		"Name": "product new",
	// 		"Price": 50,
	// 		"TaxCodeID": 3,
	// 	},
	// 	"Message": "Success"
	// }
	r.POST("/product/register", mdle.RegisterProductHandler)

	// Get list of products
	// endpoint: localhost:9090/mini-api/product/list
	// HTTP Method: GET
	// example response:
	// {
	// 	"header": {
	// 		"process_time": 0.601567,
	// 		"messages": null,
	// 		"reason": "",
	// 		"error_code": 0
	// 	},
	// 	"data": [
	// 		{
	// 			"ID": 7,
	// 			"Name": "Lucky Stretch",
	// 			"Price": 1000,
	// 			"TaxCodeID": 2,
	// 			"TaxType": "tobacco",
	// 		},
	// 		{
	// 			"ID": 8,
	// 			"Name": "Big Mac",
	// 			"Price": 1000,
	// 			"TaxCodeID": 1,
	// 			"TaxType": "food",
	// 		},
	// 		{
	// 			"ID": 9,
	// 			"Name": "Movie",
	// 			"Price": 150,
	// 			"TaxCodeID": 3,
	// 			"TaxType": "entertainment",
	// 		},
	// 		{
	// 			"ID": 10,
	// 			"Name": "product new",
	// 			"Price": 50,
	// 			"TaxCodeID": 3,
	// 			"TaxType": "entertainment",
	// 		}
	// 	],
	// 	"Message": "Success"
	// }
	r.GET("/product/list", mdle.GetListProductHandler)

	// Create transaction for user
	// endpoint: localhost:9090/mini-api/transaction/create
	// HTTP Method: POST
	//	+---------------------------------------------------------------------------------------------------+
	// 	| Params			| type		| null	| Default	| Description									|
	//	----------------------------------------------------------------------------------------------------+
	// 	| transaction_item	| string	| NO	| 			| This is list of product join with quantity	|
	// 	| 					| 			| 		| 			| U can get get list products from product/list	|
	// 	| 					| 			| 		| 			| For format is: [product_id]:[quantity]		|
	// 	| 					| 			| 		| 			| U can multiple assign separate with comma		|
	// 	| 					| 			| 		| 			| Example: 7:10,8:2,9:3							|
	// 	| 					| 			| 		| 			| 7:10 = assign "Lucky Strike" with 10 qty		|
	//	----------------------------------------------------------------------------------------------------+
	// 	| user_id			| float64	| NO	| 			| U can get information user from user/list		|
	// 	| 					| 			| 		| 			| just put user_id								|
	//	----------------------------------------------------------------------------------------------------+
	// example response:
	// {
	// 	"header": {
	// 		"process_time": 8.843066,
	// 		"messages": null,
	// 		"reason": "",
	// 		"error_code": 0
	// 	},
	// 	"data": {
	// 		"ID": 30,
	// 		"UserID": 1,
	// 		"TransactionDate": "2018-09-10T00:00:00Z",
	// 		"DetailTransaction": [
	// 			{
	// 				"ID": 86,
	// 				"Product": {
	// 					"ID": 7,
	// 					"Name": "Lucky Stretch",
	// 					"Price": 1000,
	// 					"TaxCodeID": 2,
	// 					"Quantity": 2
	// 				},
	// 				"Tax": 50,
	// 				"SubTotalPrice": 2050
	// 			},
	// 			{
	// 				"ID": 87,
	// 				"Product": {
	// 					"ID": 8,
	// 					"Name": "Big Mac",
	// 					"Price": 1000,
	// 					"TaxCodeID": 1,
	// 					"Quantity": 1
	// 				},
	// 				"Tax": 100,
	// 				"SubTotalPrice": 1100
	// 			},
	// 			{
	// 				"ID": 88,
	// 				"Product": {
	// 					"ID": 9,
	// 					"Name": "Movie",
	// 					"Price": 150,
	// 					"TaxCodeID": 3,
	// 					"Quantity": 10
	// 				},
	// 				"Tax": 0,
	// 				"SubTotalPrice": 1500
	// 			}
	// 		],
	// 		"TotalPrice": 4650
	// 	},
	// 	"Message": "Success"
	// }
	r.POST("/transaction/create", mdle.CreateTransactionHandler)

	// Get information of list transaction
	// this is not spesific detail transaction
	// just list of transaction
	// endpoint: localhost:9090/mini-api/transaction/list
	// HTTP Method: GET
	// example response:
	// {
	// "header": {
	// 	"process_time": 0.374177,
	// 	"messages": null,
	// 	"reason": "",
	// 	"error_code": 0
	// },
	// 		data": [
	// 		{
	// 			"ID": 2,
	// 			"UserID": 1,
	// 			"TransactionDate": "2018-09-10T00:00:00Z",
	// 		},
	// 		{
	// 			"ID": 3,
	// 			"UserID": 1,
	// 			"TransactionDate": "2018-09-10T00:00:00Z",
	// 		},
	// 		{
	// 			"ID": 4,
	// 			"UserID": 1,
	// 			"TransactionDate": "2018-09-10T00:00:00Z",
	// 		}
	// 	],
	//     "Message": "Success"
	// }
	r.GET("/transaction/list", mdle.GetListTransactionHandler)

	// Get information of list detail transaction
	// this is detail transaction by ID transaction
	// endpoint: localhost:9090/mini-api/transaction/detail?id=[id_transaction]
	// id_transaction u can get from /transaction/list
	// HTTP Method: GET
	// example response:
	// {
	// 	"header": {
	// 		"process_time": 2.0947829999999996,
	// 		"messages": null,
	// 		"reason": "",
	// 		"error_code": 0
	// 	},
	// 	"data": {
	// 		"id": 31,
	// 		"user_id": 1,
	// 		"transaction_date": "2018-09-10T22:52:50.796346Z",
	// 		"detail_transaction": [
	// 			{
	// 				"transaction_item_id": 89,
	// 				"Product": {
	// 					"id": 7,
	// 					"name": "Lucky Stretch",
	// 					"price": 1000,
	// 					"tax_code": 2,
	// 					"tax_type": "tobacco",
	// 					"quantity": 2
	// 				},
	// 				"tax": 50,
	// 				"sub_total_price": 2050
	// 			},
	// 			{
	// 				"transaction_item_id": 90,
	// 				"Product": {
	// 					"id": 8,
	// 					"name": "Big Mac",
	// 					"price": 1000,
	// 					"tax_code": 1,
	// 					"tax_type": "food",
	// 					"quantity": 1
	// 				},
	// 				"tax": 100,
	// 				"sub_total_price": 1100
	// 			},
	// 			{
	// 				"transaction_item_id": 91,
	// 				"Product": {
	// 					"id": 9,
	// 					"name": "Movie",
	// 					"price": 150,
	// 					"tax_code": 3,
	// 					"tax_type": "entertainment",
	// 					"quantity": 10
	// 				},
	// 				"tax": 0,
	// 				"sub_total_price": 1500
	// 			}
	// 		],
	// 		"total_price": 4650
	// 	},
	// 	"Message": "Success"
	// }
	r.GET("/transaction/detail", mdle.GetDetailTransactionHandler)
}
