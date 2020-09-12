package models

// Character struct ...
type Character struct {
	ID       int    `json:"id" gorm:"primary_key;size:25"`
	UnitID   int    `json:"unitId"`
	UnitName string `json:"unitName"`
	Icon     string `json:"icon"`
	Position int    `json:"position"`
}

func (Character) TableName() string {
	return "characters"
}
