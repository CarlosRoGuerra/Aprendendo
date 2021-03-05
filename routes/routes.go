package routes

import (
	"net/http"

	"github.com/CarlosRoGuerra/Aprendendo/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)

}
