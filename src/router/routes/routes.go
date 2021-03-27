package routes

import (
	"api/src/middlewares"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Route represeta todas as rotas do sistema
type Route struct {
	Url          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

//Config configura todas as rotas do router
func Config(r *mux.Router) *mux.Router {
	routes := accountRoutes
	routes = append(routes, loginRoute, transferRoutes[1], transferRoutes[0])
	for _, route := range routes {

		if route.AuthRequired {
			r.HandleFunc(route.Url, middlewares.Logger(middlewares.Auth(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.Url, route.Function).Methods(route.Method)
		}
	}
	fmt.Println("rotas configuradas")
	return r
}
