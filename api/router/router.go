package router

import (
	"online-store/api/router/routes"

	"github.com/gorilla/mux"
)

type RouterInit struct {
	Router *mux.Router
}

func (r *RouterInit) Init() {
	r.Router = mux.NewRouter().StrictSlash(true)
	r.Router = routes.LoadRoute(r.Router)
}
