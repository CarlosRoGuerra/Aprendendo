package controllers

import (
	"net/http"
	"text/template"

	"github.com/CarlosRoGuerra/Aprendendo/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsjogos := models.BuscaTodosOsjogos()
	temp.ExecuteTemplate(w, "Index", todosOsjogos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}
