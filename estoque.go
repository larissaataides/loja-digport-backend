package main

import (
	"errors"

	"github.com/larissaataides/loja-digport-backend/model"
)

var ListaDeProdutos []model.Produto = []model.Produto{}

func criarEstoque() { //retorno da nossa função, retorna um array e do tipo produto
	produtos := []model.Produto{
		{
			Id:                  "1",
			Nome:                "Produto diversificado",
			Descricao:           "Descrição do Produto diversificado",
			Categoria:           "Planta",
			Valor:               10.99,
			QuantidadeEmEstoque: 3,
			Imagem:              "planta.png",
		},

		{

			Id:                  "2",
			Nome:                "Produto diferenciado",
			Descricao:           "Descrição do Produto diferenciado",
			Categoria:           "caneca",
			Valor:               350.40,
			QuantidadeEmEstoque: 5,
			Imagem:              "caneca.png",
		},

		{

			Id:                  "3",
			Nome:                "Produto qualificado",
			Descricao:           "Descrição do Produto qualificado",
			Categoria:           "livro",
			Valor:               63.30,
			QuantidadeEmEstoque: 20,
			Imagem:              "livro.png",
		},
	}
	ListaDeProdutos = produtos
}

/* func filtrarProdutos(nome string) []model.Produto {
	//criar uma função para buscar produtos pelo tipo dele ex: http/produtos?nome

	produtos := []model.Produto{
		{
			Id:                  "2",
			Nome:                "Produto diferenciado",
			Descricao:           "Descrição do Produto diferenciado",
			Categoria:           "caneca",
			Valor:               350.40,
			QuantidadeEmEstoque: 5,
			Imagem:              "caneca.png",
		},
	}
	return produtos
}
*/

func buscaPorNome(nome string) []model.Produto {

	resultado := []model.Produto{}

	for _, produto := range ListaDeProdutos {
		if produto.Nome == nome {
			resultado = append(resultado, produto)
		}
	}
	return resultado
}

func adicionarProduto(produtoNovo model.Produto) error {
	for _, produtoDaLista := range ListaDeProdutos {
		if produtoNovo.Id == produtoDaLista.Id {
			return errors.New("produto com ID já cadastrada")
		}
	}

	ListaDeProdutos = append(ListaDeProdutos, produtoNovo)
	return nil

}
