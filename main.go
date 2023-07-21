package main

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html")) //var para indicar o caminho das Templates//

func main() { // criar um server para aplicação web na porta 8000 //
	http.HandleFunc("/", index) // "/" - caminho que voce quer atender, Index - quem vai atender ? No caso uma func Index irá atender//
	http.ListenAndServe(":8000", nil)

}
func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
} // ResponsWriter - quem vai ler a requisição, Request - A requisição //
