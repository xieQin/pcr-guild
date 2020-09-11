package models

// Character struct ...
type Character struct {
	ID       int    `json:"id"`
	UnitID   int    `json:"unitId"`
	UnitName string `json:"unitName"`
	Icon     string `json:"icon"`
	Position int    `json:"position"`
}
