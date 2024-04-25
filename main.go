package main

import (
	"fmt"
)

func main() {
	//prod := criarEstoque()                                  //para chamar a função de criarEstoque
	//fmt.Println("Esse é o catálogo da loja xxx: %+v", prod) //Para imprimir a lista da função CriarEstoque
	//StartServer()

	nomeProduto := ""
	produtosFiltrados := filtrarProdutos(nomeProduto)

	fmt.Print(produtosFiltrados)

}
