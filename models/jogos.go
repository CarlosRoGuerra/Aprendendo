package models

import (
	"github.com/CarlosRoGuerra/Go_Wep_App-Cr/db"
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

func DeletaJogo(id string) {
	db := db.ConectaComBancoDeDados()

	deletarOJogo, err := db.Prepare("delete from jogos where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deletarOJogo.Exec(id)
	defer db.Close()
}

func EditaJogo(id string) Jogo {
	db := db.ConectaComBancoDeDados()

	jogoDoBanco, err := db.Query("select from jogos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	jogoParaAtualizar := Jogo{}

	for jogoDoBanco.Next() {
		var id int
		var nome, genero, plataforma string
		var preco float64

		err = jogoDoBanco.Scan(&id, &nome, &genero, &preco, &plataforma)
		if err != nil {
			panic(err.Error())
		}
		jogoParaAtualizar.Id = id
		jogoParaAtualizar.Nome = nome
		jogoParaAtualizar.Genero = genero
		jogoParaAtualizar.Preco = preco
		jogoParaAtualizar.Plataforma = plataforma
	}
	defer db.Close()
	return jogoParaAtualizar
}

func AtualizaJogo(id int, nome, plataforma, genero string, preco float64) {
	db := db.ConectaComBancoDeDados()

	AtualizaJogo, err := db.Prepare("update jogos set nome=$1, genero=$2, preco=$3, plataforma=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}
	AtualizaJogo.Exec(nome, genero, preco, plataforma, id)
	defer db.Close()
}
