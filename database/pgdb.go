package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBp  *gorm.DB
	errr error
)

func ConectaDBp() {

	dbConect := "host=mouse.db.elephantsql.com user=lgngbpdo password=5NqqDGWFkazYmraBjtPhJAeUoJzFJ5Kt dbname=lgngbpdo port=5432 sslmode=disable"
	DBp, errr = gorm.Open(postgres.Open(dbConect), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com o Banco de dados")
	}
}

/*
func ConectaDB() {
	dbConect := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dbConect), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com o Banco de dados")
	}
}
*/
