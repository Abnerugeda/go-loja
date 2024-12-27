package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/Abnerugeda/go-loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.FindProducts()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversao do preço: ", err.Error())
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversao da quantidae: ", err.Error())
		}

		produto := models.Produto{
			Nome:       nome,
			Descricao:  descricao,
			Preco:      precoConvertido,
			Quantidade: quantidadeConvertida,
		}
		models.InsertProdutos(produto)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")

	models.DeletarProduto(idProduto)

	http.Redirect(w, r, "/", http.StatusFound)
}

func EditView(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.FindOneProduct(idProduto)

	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversao do preço: ", err.Error())
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversao da quantidae: ", err.Error())
		}

		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversao da quantidae: ", err.Error())
		}

		produto := models.Produto{
			Id:         idConvertido,
			Nome:       nome,
			Descricao:  descricao,
			Preco:      precoConvertido,
			Quantidade: quantidadeConvertida,
		}
		models.UpdateProduto(produto)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
