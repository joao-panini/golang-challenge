package routes

import (
	"api/src/controllers"
	"net/http"
)

var transferRoutes = []Route{
	{
		Url:          "/transfers",
		Method:       http.MethodGet,
		Function:     controllers.GetAllCurrentAcctTransactions,
		AuthRequired: true,
	},
	{
		Url:          "/transfers",
		Method:       http.MethodPut,
		Function:     controllers.CreateTransaction,
		AuthRequired: true,
	},
}
