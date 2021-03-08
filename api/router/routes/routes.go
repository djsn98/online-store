package routes

import "github.com/gorilla/mux"

func LoadRoute(r *mux.Router) *mux.Router {
	userRoute(r.PathPrefix("/user").Subrouter())
	productRoute(r.PathPrefix("/product").Subrouter())
	orderRoute(r.PathPrefix("/order").Subrouter())

	return r
}
