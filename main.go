package main

import (
	"net/http"

	"github.com/CarlosRoGuerra/Aprendendo/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)

}
