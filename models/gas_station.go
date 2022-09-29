package models

type Gas_Station struct {
	ID         uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name       string `json:"name"`
	CategoryID uint32 `json:"category_id"`
}

var Gas_Stations []Gas_Station

type Category struct {
	ID           uint32 `gorm:"primary_key;auto_increment" json:"id"`
	CategoryName string `json:"category_name"`
}

var Categories []Category

type GS_Category struct {
	Name         string `json:"gas_station_name"`
	CategoryID   uint32 `json:"category_id"`
	CategoryName string `json:"category_name"`
}
