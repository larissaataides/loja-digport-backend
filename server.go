/*
	func StartServer() {
		http.HandleFunc("/produtos", produtosHandler)

		http.ListenAndServe(":8080", nil)
	}

/*

	func produtosHandler(w http.ResponseWriter, r *http.Request) {
		produtos := criarEstoque()

		json.NewEncoder(w).Encode(produtos) //encoder é responsável por pegar um objeto e traduzir para o Json

}

	func produtosHandler(w http.ResponseWriter, r *http.Request) {
		queryNome := r.URL.Query().Get("nome")
		if queryNome != "" {
			produtosPorNome := buscaPorNome(queryNome)
			json.NewEncoder(w).Encode(produtosPorNome)
		} else {
			produtos := ListaDeProdutos
			json.NewEncoder(w).Encode(produtos)
		}
	}

	func addProduto(w http.ResponseWriter, r *http.Request) {
		var produto model.Produto
		json.NewDecoder(r.Body).Decode(&produto)
		criarEstoque()
		w.WriteHeader(http.StatusCreated)
	}
*/
package main

import "github.com/larissaataides/loja-digport-backend/routes"

func StartServer() {
	routes.HandleRequests()
}
