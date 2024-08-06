package models

import (
	"log"

	"github.com/marin/api/db"
)

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func BuscaTodosProdutos() []Produto {
	db := db.ConectaComBanco()
	selectProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}
	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriaNovoProduto(nome string, descicao string, preco float64, quantidade int) {
	db := db.ConectaComBanco()
	insertDados, err := db.Prepare("insert into produtos(nome,descricao,preco,quantidade) values($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}
	insertDados.Exec(nome, descicao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(idtProduto string) {
	db := db.ConectaComBanco()
	deletaProdutoPorId, err := db.Prepare("delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}
	log.Println("[INFO] - DeletaProduto - Produto com id", idtProduto, "deletado")
	deletaProdutoPorId.Exec(idtProduto)
	defer db.Close()
}

func BuscaPorProduto(idtProduto string) Produto {
	db := db.ConectaComBanco()
	produtoBuscadoPorId, err := db.Query("select * from produtos where id = $1", idtProduto)
	if err != nil {
		panic(err.Error())
	}
	produtoParaAtualizar := Produto{}
	for produtoBuscadoPorId.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err = produtoBuscadoPorId.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id string, nome string, descicao string, preco float64, quantidade int) {
	db := db.ConectaComBanco()
	insertDados, err := db.Prepare("UPDATE produtos SET nome = $2, descricao = $3, preco = $4, quantidade = $5 WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}
	insertDados.Exec(id, nome, descicao, preco, quantidade)
	defer db.Close()
}
