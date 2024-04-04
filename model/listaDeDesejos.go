package model

type ListaDeDesejos struct {
	Id        string
	UserId    string
	ProdutoId []string //um array de para adicionar mais produtos e não somente um
}
