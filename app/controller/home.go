package controller

import (

	// m "golang-study/app/model"

	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Hello(message string) {
	var name string

	fmt.Println(message)
	fmt.Scanf("%s", &name)
	fmt.Printf("Hi, %v", name)
}

// func Marshal(v interface{}) ([]byte, error)

// Home - controller home page
func Home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	fmt.Fprintln(w, "tets")

	_message := "Hello world!"

	go Hello(_message)


	// profile := m.Home{name: "Hello world!"} // задаю данные для json

	// js, err := json.Marshal(profile) // обработка ошибок обязательна для json

	// if err != nil {
	// 	panic(err)
	// }

	// w.Header().Set("Content-Type", "application/json") // Задаю формат вывода
	// w.Write(js)
}
