package routes

import (
	"net/http"

	"github.com/CarlosRoGuerra/Go_Wep_App-Cr/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)

}
