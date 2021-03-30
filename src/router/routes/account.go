package routes

import (
	"api/src/controllers"
	"net/http"
)

var accountRoutes = []Route{
	{
		Url:          "/accounts",
		Method:       http.MethodPost,
		Function:     controllers.CreateAccount,
		AuthRequired: false,
	},
	{
		Url:          "/accounts",
		Method:       http.MethodGet,
		Function:     controllers.GetAccounts,
		AuthRequired: true,
	},
	{
		Url:          "/account/{accountID}/balance",
		Method:       http.MethodGet,
		Function:     controllers.GetAccountBalanceById,
		AuthRequired: false,
	},
}
