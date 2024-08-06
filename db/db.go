package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConectaBancoDados() *sql.DB {
	connStr := "user=postgres dbname=digport_loja password=digport host=localhost sslmode=disable" //- não é seguro ter a senha aqui
	//dbPass := os.Getenv("DB_PASS") //no terminal DB_Pass=senha para definir a senha local, uma vez terminal fechado, senha some. comando echo $DB_Pass mostra a senha
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db

}
