package main

import (
	"fmt"
	"net/http"
	"online-store/api/router"
)

func main() {
	router := router.InitRouter()
	http.Handle("/", router)

	var address = "localhost:9000"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
