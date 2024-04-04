package model

type Carrinho struct {
	IdUsuario       string
	IdCarrinho      string
	QuantidadeTotal []int
	ValorTotal      float64
	InfosProduto    []InfosProduto
}

type InfosProduto struct {
	ProdutoId           string
	QuantidadeDeProduto []int
}
