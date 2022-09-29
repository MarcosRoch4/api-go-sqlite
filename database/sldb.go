package database

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaDB() {

	dbConect := "./cheapgas.db"
	DBp, errr = gorm.Open(sqlite.Open(dbConect), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com o Banco de dados")
	}
}
