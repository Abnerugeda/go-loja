package models

import (
	"database/sql"
	"fmt"

	"github.com/Abnerugeda/go-loja/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func FindProducts() []Produto {
	db := db.ConnectDB()
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
	defer db.Close()
	return produtos
}

func InsertProdutos(produto Produto) sql.Result {
	db := db.ConnectDB()

	res, err := db.Exec(fmt.Sprintf(`INSERT INTO produtos (nome, descricao, preco, quantidade)
						VAlUES ('%s', '%s', %f, %d)`, produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	return res
}
