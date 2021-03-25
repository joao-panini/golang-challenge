package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

//retorna as rotas configuradas
func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
