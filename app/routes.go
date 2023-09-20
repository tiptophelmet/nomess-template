package app

import (
	"github.com/tiptophelmet/nomess-core/v5/router"
	"github.com/tiptophelmet/nomess-template/handler"
)

func initRoutes() {
	router.Init()
	router.Handle("/helloworld", handler.HelloWorld).Methods("GET")

	router.Handle("/item", handler.CreateItem).Methods("POST")
	router.Handle("/item/{id}", handler.GetItem).Methods("GET")
	router.Handle("/item", handler.GetItems).Methods("GET")
	router.Handle("/item/{id}", handler.UpdateItem).Methods("PUT")
	router.Handle("/item/{id}", handler.DeleteItem).Methods("DELETE")
}
