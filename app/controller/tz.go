package controller

import (

	// m "gggg/app/model"

	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var a = 0

// Tz - count modulo
func Tz(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	a++
	if a%5 == 0 {
		log.Print(a)
	}
	fmt.Fprintln(w, "tz")
}
