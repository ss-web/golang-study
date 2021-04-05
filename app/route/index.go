package route

import (
	c "golang-study/app/controller"

	"github.com/julienschmidt/httprouter"
)

var Router *httprouter.Router

func init() {
	Router = httprouter.New()

	Router.GET("/", c.Home)
	Router.GET("/tz", c.Tz)
}
