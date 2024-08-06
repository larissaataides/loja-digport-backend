package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/larissaataides/loja-digport-backend/controller"
)

func HandleRequests() {
	route := mux.NewRouter()
	route.HandleFunc("/produtos", controller.BuscaProdutosHandler).Methods("GET")
	route.HandleFunc("/produto", controller.BuscarProdutosPorNomeHandler).Methods("GET")
	route.HandleFunc("/produtos", controller.CriarProdutosHandler).Methods("POST")

	http.ListenAndServe(":8080", route)

}
