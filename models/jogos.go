package models

import (
	"github.com/CarlosRoGuerra/Aprendendo/db"
)

type Jogo struct {
	id     int
	Nome   string
	Genero string
	Preco  float64

	Plataforma string
}

func BuscaTodosOsjogos() []Jogo {

	db := db.ConectaComBancoDeDados()

	selectDeTodosOsJogos, err := db.Query("select *from jogos")
	if err != nil {
		panic(err.Error())
	}
	p := Jogo{}
	jogos := []Jogo{}

	for selectDeTodosOsJogos.Next() {
		var id int
		var nome, genero, plataforma string
		var preco float64

		err = selectDeTodosOsJogos.Scan(&id, &nome, &genero, &plataforma, &preco)
		if err != nil {
			panic(err.Error())
		}
		p.Nome = nome
		p.Genero = genero
		p.Plataforma = plataforma
		p.Preco = preco

		jogos = append(jogos, p)
	}
	defer db.Close()
	return jogos
}
