package main

import (
	"encoding/json"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/produtos", produtosHandler)

	http.ListenAndServe(":8080", nil)
}

func produtosHandler(w http.ResponseWriter, r *http.Request) {
	produtos := criarEstoque()

	json.NewEncoder(w).Encode(produtos) //encoder é responsável por pegar um objeto e traduzir para o Json

}
