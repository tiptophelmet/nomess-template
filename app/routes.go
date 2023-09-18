package app

import (
	"fmt"

	"github.com/tiptophelmet/nomess-core/v3/router"
	"github.com/tiptophelmet/nomess-template/handler"
)

func initRoutes() {
	router.Init()
	router.Handle("/helloworld", handler.HelloWorld).Methods("GET")
	router.Handle("/item", handler.CreateItem).Methods("POST")

	fmt.Println("Routes init OK!")
}
