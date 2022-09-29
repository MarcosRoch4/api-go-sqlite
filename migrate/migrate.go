package migrate

import (
	"github.com/MarcosRoch4/api-go-sqlite/database"
	"github.com/MarcosRoch4/api-go-sqlite/models"
)

func IniMigrate() {
	//database.ConectaDB()
	database.ConectaDB()
	database.DB.AutoMigrate(&models.Fuel_Value{}, &models.Fuel{}, &models.Gas_Station{}, &models.Category{})
}
