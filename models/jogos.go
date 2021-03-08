package models

import (
	"github.com/CarlosRoGuerra/Aprendendo/db"
)

type Jogo struct {
	Id     int
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
		p.Id = id
		p.Nome = nome
		p.Genero = genero
		p.Plataforma = plataforma
		p.Preco = preco

		jogos = append(jogos, p)
	}
	defer db.Close()
	return jogos
}

func CriarNovoJogo(nome, genero string, preco float64, plataforma string) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into jogos(nome, genero, preco, plataforma) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, genero, preco, plataforma)

	defer db.Close()
}
