package models

type Fuel struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	FuelName string `json:"fuel_name"`
}

var Fuels []Fuel

type Fuel_Value struct {
	ID             uint32  `gorm:"primary_key;auto_increment" json:"id"`
	Price          float64 `json:"price"`
	Fuel_Id        uint32  `gorm:"not null" json:"fuel_id"`
	Gas_Station_Id uint32  `gorm:"not null" json:"gas_station_id"`
}

var Fuel_Values []Fuel_Value

type GS_Fuel struct {
	Gas_Station_Id uint32
	Name           string
	FuelName       string
	Price          float64
}
