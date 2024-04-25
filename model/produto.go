package model

type Produto struct {
	Id                  string
	Nome                string
	Descricao           string
	Categoria           string
	Valor               float64
	QuantidadeEmEstoque int
	Imagem              string
}
