package controller

import (

	// m "golang-study/app/model"

	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Home - controller home page
func Home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	fmt.Fprintln(w, "tets")

	// profile := m.Home{name: "Hello world!"} // задаю данные для json

	// js, err := json.Marshal(profile) // обработка ошибок обязательна для json

	// if err != nil {
	// 	panic(err)
	// }

	// w.Header().Set("Content-Type", "application/json") // Задаю формат вывода
	// w.Write(js)
}
