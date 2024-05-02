package model

type Produto struct {
	Id                  string  `json:"id"`
	Nome                string  `json:"nome"`
	Descricao           string  `json:"descricao"`
	Categoria           string  `json:"categoria"`
	Valor               float64 `json:"valor"`
	QuantidadeEmEstoque int     `json:"quantidadeEmEstoque"`
	Imagem              string  `json:"imagem"`
}
