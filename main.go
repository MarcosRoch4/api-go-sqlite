package main

import (
	"fmt"

	"github.com/MarcosRoch4/api-rest-go/migrate"
	"github.com/MarcosRoch4/api-rest-go/routes"
)

func main() {

	migrate.IniMigrate()
	fmt.Println("Iniciando o servidor Rest com GO!")
	routes.HandleRequest()
}
