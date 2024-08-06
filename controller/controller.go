package controller

import (
	"encoding/json"
	"net/http"

	"github.com/larissaataides/loja-digport-backend/model"
)

func BuscaProdutosHandler(w http.ResponseWriter, r *http.Request) {
	produtos := model.BuscaTodosProdutos()
	json.NewEncoder(w).Encode(produtos) //passar uma lista de produtos em json pra chamar nossa API
}

func BuscarProdutosPorNomeHandler(w http.ResponseWriter, r *http.Request) {

}

func CriarProdutosHandler(w http.ResponseWriter, r *http.Request) {

}
