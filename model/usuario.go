package model

type Usuario struct {
	Nome  string
	Id    string
	Email string
	senha string //a senha n precisa ser pública, por isso está em minúsculo
}

//pega a senha e faz um hash - crypt biblioteca

//valida a senha
