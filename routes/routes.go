package routes

import (
	"net/http"

	"github.com/Abnerugeda/go-loja/controllers"
)

func GetRoutes() {
	http.HandleFunc("/", controllers.Index)
}
