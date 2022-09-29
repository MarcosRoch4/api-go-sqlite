package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBm  *gorm.DB
	errm error
)

func ConectaMSDB() {

	dbConect := "rfwlklxwsunzqv9qequq:pscale_pw_Cn02G8PfIDiU8Jsofu7bLTwEO82zXhhvtf1XXw0YhX0@tcp(us-east.connect.psdb.cloud)/cheapgas?tls=true"
	DBm, errm = gorm.Open(mysql.Open(dbConect), &gorm.Config{})

	if errm != nil {
		log.Panic("Erro ao conectar com o Banco de dados")
	}

}
