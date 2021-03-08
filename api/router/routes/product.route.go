package routes

import "github.com/gorilla/mux"

func productRoute(r *mux.Router) {
	r.HandleFunc("/read", handlerIndex)
}
