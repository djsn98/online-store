package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var message = "Home"
	w.Write([]byte(message))
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	var message = "Products"
	w.Write([]byte(message))
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	var message = "Articles"
	w.Write([]byte(message))
}

func InitRouter() *mux.Router {
	Router.HandleFunc("/", HomeHandler)
	Router.HandleFunc("/products", ProductsHandler)
	Router.HandleFunc("/articles", ArticlesHandler)

	return Router
}
