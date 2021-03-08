package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/CarlosRoGuerra/Go_Wep_App-Cr/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsjogos := models.BuscaTodosOsjogos()
	temp.ExecuteTemplate(w, "Index", todosOsjogos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		genero := r.FormValue("genero")
		preco := r.FormValue("preco")
		plataforma := r.FormValue("plataforma")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversao do preco", err.Error())
		}

		models.CriarNovoJogo(nome, genero, precoConvertido, plataforma)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoJogo := r.URL.Query().Get("id")
	models.DeletaJogo(idDoJogo)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoJogo := r.URL.Query().Get("id")
	jogo := models.EditaJogo(idDoJogo)
	temp.ExecuteTemplate(w, "Edit", jogo)

}
