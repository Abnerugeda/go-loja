package main

import (
	"html/template"
	"net/http"

	"github.com/Abnerugeda/go-loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := models.FindProducts()
	temp.ExecuteTemplate(w, "Index", produtos)
}
