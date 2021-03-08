package routes

import "github.com/gorilla/mux"

func userRoute(r *mux.Router) {
	r.HandleFunc("/login", handlerIndex)
	r.HandleFunc("/register", handlerIndex)
}
