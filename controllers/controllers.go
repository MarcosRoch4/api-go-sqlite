package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MarcosRoch4/api-go-sqlite/database"
	"github.com/MarcosRoch4/api-go-sqlite/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

// Postos de Combutíveis Disponíveis

func TodosPosto(w http.ResponseWriter, r *http.Request) {
	var p []models.Gas_Station
	database.DB.Order("id").Table("gas_stations").Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmPosto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var posto models.Gas_Station
	database.DB.Table("gas_stations").First(&posto, id)
	json.NewEncoder(w).Encode(posto)

}

func NovoPosto(w http.ResponseWriter, r *http.Request) {
	var posto models.Gas_Station
	json.NewDecoder(r.Body).Decode(&posto)
	database.DB.Table("gas_stations").Create(&posto)
	json.NewEncoder(w).Encode(&posto)
}

func DeletaPosto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var posto models.Gas_Station
	database.DB.Table("gas_stations").Delete(&posto, id)
	json.NewEncoder(w).Encode(posto)
}

func EditaPosto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var posto models.Gas_Station
	database.DB.Table("gas_stations").First(&posto, id)
	json.NewDecoder(r.Body).Decode(&posto)
	database.DB.Table("gas_stations").Save(&posto)
	json.NewEncoder(w).Encode(posto)
}

// Combutíveis Disponíveis

func TodosCombustiveis(w http.ResponseWriter, r *http.Request) {
	var p []models.Fuel
	database.DB.Order("id").Table("fuels").Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmCombustivel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var combustivel models.Fuel
	database.DB.Table("fuels").First(&combustivel, id)
	json.NewEncoder(w).Encode(combustivel)

}

func NovoCombustivel(w http.ResponseWriter, r *http.Request) {
	var combustivel models.Fuel
	json.NewDecoder(r.Body).Decode(&combustivel)
	database.DB.Table("fuels").Create(&combustivel)
	json.NewEncoder(w).Encode(&combustivel)
}

func DeletaCombustivel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var combustivel models.Fuel
	database.DB.Table("fuels").Delete(&combustivel, id)
	json.NewEncoder(w).Encode(combustivel)
}

func EditaCombustivel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var combustivel models.Fuel
	database.DB.Table("fuels").First(&combustivel, id)
	json.NewDecoder(r.Body).Decode(&combustivel)
	database.DB.Table("fuels").Save(&combustivel)
	json.NewEncoder(w).Encode(combustivel)
}

// Valores dos Combustíveis

func TodosValores(w http.ResponseWriter, r *http.Request) {
	var p []models.Fuel_Value
	database.DB.Order("id").Table("fuel_values").Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmValor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var valor models.Fuel_Value
	database.DB.Table("fuel_values").First(&valor, id)
	json.NewEncoder(w).Encode(valor)

}

func NovoValor(w http.ResponseWriter, r *http.Request) {
	var valor models.Fuel_Value
	json.NewDecoder(r.Body).Decode(&valor)
	database.DB.Table("fuel_values").Create(&valor)
	json.NewEncoder(w).Encode(&valor)
}

func DeletaValor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var valor models.Fuel_Value
	database.DB.Table("fuel_values").Delete(&valor, id)
	json.NewEncoder(w).Encode(valor)
}

func EditaValor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var valor models.Fuel_Value
	database.DB.Table("fuel_values").First(&valor, id)
	json.NewDecoder(r.Body).Decode(&valor)
	database.DB.Table("fuel_values").Save(&valor)
	json.NewEncoder(w).Encode(valor)
}

// Categoria dos Postos de Combustíveis

func TodasCategorias(w http.ResponseWriter, r *http.Request) {

	//	database.DB.Debug().AutoMigrate(&models.Gas_Station{}, &models.Category{})
	var p []models.Category
	database.DB.Order("id").Table("categories").Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaCategoria(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var categoria models.Category
	database.DB.Table("categories").First(&categoria, id)
	json.NewEncoder(w).Encode(categoria)

}

func NovaCategoria(w http.ResponseWriter, r *http.Request) {
	var categoria models.Category
	json.NewDecoder(r.Body).Decode(&categoria)
	database.DB.Table("categories").Create(&categoria)
	json.NewEncoder(w).Encode(&categoria)
}

func DeletaCategoria(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var categoria models.Category
	database.DB.Table("categories").Delete(&categoria, id)
	json.NewEncoder(w).Encode(categoria)
}

func EditaCategoria(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var categoria models.Category
	database.DB.Table("categories").First(&categoria, id)
	json.NewDecoder(r.Body).Decode(&categoria)
	database.DB.Table("categories").Save(&categoria)
	json.NewEncoder(w).Encode(categoria)
}

// Postos de Combustíveis Ordenados

func PostoPorCategoria(w http.ResponseWriter, r *http.Request) {
	gs_category := []models.GS_Category{}

	database.DB.Raw("SELECT gas_stations.name, gas_stations.category_id, categories.category_name FROM gas_stations " +
		"LEFT JOIN categories ON categories.ID = gas_stations.category_id ORDER BY gas_stations.category_id DESC").
		Scan(&gs_category)

	fmt.Println("Resultado ===>", gs_category)

	json.NewEncoder(w).Encode(gs_category)

}

func PostoPorCombustivel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fuel_id := vars["id"]

	gs_fuel := []models.GS_Fuel{}

	database.DB.Raw("SELECT fv.gas_station_id,gs.name, f.fuel_name,fv.price "+
		"FROM fuel_values fv "+
		"LEFT JOIN gas_stations gs ON gs.id = fv.gas_station_id "+
		"LEFT JOIN fuels f ON f.id = fv.fuel_id "+
		"WHERE f.id = ? "+
		"ORDER BY fv.price ASC", fuel_id).
		Scan(&gs_fuel)

	json.NewEncoder(w).Encode(gs_fuel)

}

func Resultado(w http.ResponseWriter, r *http.Request) {

	gs_category := []models.GS_Category{}

	database.DB.Raw("SELECT gas_stations.name, gas_stations.category_id, categories.category_name FROM gas_stations " +
		"LEFT JOIN categories ON categories.ID = gas_stations.category_id ORDER BY gas_stations.category_id DESC").
		Scan(&gs_category)

	fmt.Println("Resultado ===>", gs_category)

	json.NewEncoder(w).Encode(gs_category)

}
