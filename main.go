package main

import (
	"online-store/api/router"
	"online-store/api/server"
)

func main() {
	//Membuat router dan inisialisai router
	router := router.RouterInit{}
	router.Init()

	//Config
	host := "localhost:"
	port := "9000"

	//Membuat server, inisialisasi server, dan memasang router ke server
	server := server.Server{}
	server.Init(host+port, router.Router)

	//Menyalakan Server
	server.Run()
}
