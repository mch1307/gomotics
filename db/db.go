package db

// GenericItem holds definition of item equipment
type GenericItem struct {
	id        int
	provider  string
	name      string
	location  string
	value     int
	itemType  string
	switchCmd string
}

// NhcLocation holds NHC location
type NhcLocation struct {
	id int
	name string
}

// NhcEquipment holds NHC equipment
type NhcEquipment struct {
	id       int
	name     string
	location string
}
