package main

import (
	// c "gggg/app/controller"
	// db "gggg/app"
	"gggg/app/route"
	"log"
	"net/http"
)

func main() {
	//
	// db.Connect(1)
	// log.Print("запущено")
	err := http.ListenAndServe(":8081", route.Router)
	log.Fatal(err)
}
