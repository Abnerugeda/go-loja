package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/go")

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	allProducts, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for allProducts.Next() {

		err = allProducts.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)

		if err != nil {
			panic(err.Error())
		}

		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}

func insertProduct(produto Produto) sql.Result {
	var res sql.Result

	res, _ = db.Exec(fmt.Sprintf(`INSERT INTO produtos (nome,descricao,preco,quantidade) VALUES
		('%s','%s',%f,%d)`, produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade))

	return res
}
