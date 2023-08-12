package main

import (
	"github.com/tiptophelmet/nomess/app"
	"github.com/tiptophelmet/nomess/handler"
)

func main() {
	app.Handle("/api/auth/register", handler.Register)
}
