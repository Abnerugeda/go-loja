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

func FindOneProduct(id string) Produto {
	db := db.ConnectDB()

	productQuery, err := db.Prepare("SELECT * FROM produtos WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}

	query := productQuery.QueryRow(id)

	var produto Produto

	query.Scan(&produto.Id, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	return produto
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

func UpdateProduto(produto Produto) {
	db := db.ConnectDB()

	query, err := db.Prepare(`UPDATE produtos SET nome = ?, descricao = ?, preco = ?, quantidade = ? WHERE id = ?`)

	if err != nil {
		panic(err.Error())
	}
	query.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade, produto.Id)

	defer db.Close()
}

func DeletarProduto(idProduto string) {
	db := db.ConnectDB()

	delete, err := db.Prepare("DELETE FROM produtos WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	delete.Exec(idProduto)

	defer db.Close()
}
