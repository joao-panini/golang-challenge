package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRoute = Route{
	Url:          "/login",
	Method:       http.MethodPost,
	Function:     controllers.Login,
	AuthRequired: false,
}
