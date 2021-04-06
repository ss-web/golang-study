package main

import (
	// c "golang-study/app/controller"
	// db "golang-study/app"
	"golang-study/app/route"
	"log"
	"net/http"
)

func main() {
	//
	// db.Connect(1)

	log.Print("запущено")
	err := http.ListenAndServe(":8081", route.Router)
	log.Fatal(err)
}
