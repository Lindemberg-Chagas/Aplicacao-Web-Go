package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func conectaBancoDeDados() *sql.DB {
	conexao := "user=postgres, dbname=db_loja, senha=310885, host=localhost, sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := conectaBancoDeDados()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}
func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Camisa", "Bem Colorida", 59.90, 5},
		{"Meia", "Branca", 19.90, 10},
		{"Sapato", "Social Preto", 30.00, 2},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
