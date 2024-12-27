package main

import (
	"net/http"

	"github.com/Abnerugeda/go-loja/routes"
)

func main() {
	routes.GetRoutes()
	http.ListenAndServe(":8000", nil)
}
