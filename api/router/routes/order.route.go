package routes

import "github.com/gorilla/mux"

func orderRoute(r *mux.Router) {
	r.HandleFunc("/create", handlerIndex)
}
