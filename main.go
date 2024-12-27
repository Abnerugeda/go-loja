package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	_, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul bonita", Preco: 1000, Quantidade: 10},
		{"Tenis", "confortavel", 100, 1},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
