package models

// Box struct ...
type Box struct {
	ID                    int    `json:"id" gorm:"primary_key"`
	UserID                int    `json:"user_id"`
	Server                string `json:"server"`
	UnitID                int    `json:"unit_id"`
	UnitName              string `json:"unit_name"`
	Trace                 bool   `json:"trace"`
	Level                 int    `json:"level"`
	Rarity                int    `json:"rarity"`
	Promotion             int    `json:"promotion"`
	LoveLevel             int    `json:"love_level"`
	Pieces                int    `json:"pieces"`
	Slot1                 bool   `json:"slot1"`
	Slot2                 bool   `json:"slot2"`
	Slot3                 bool   `json:"slot3"`
	Slot4                 bool   `json:"slot4"`
	Slot5                 bool   `json:"slot5"`
	Slot6                 bool   `json:"slot6"`
	UniqueEquipRank       int    `json:"unique_equip_rank"`
	Icon                  string `json:"icon"`
	TargetRarity          int    `json:"target_rarity"`
	TargetPromotion       int    `json:"target_promotion"`
	TargetLoveLevel       int    `json:"target_love_level"`
	TargetUniqueEquipRank int    `json:"target_unique_equip_rank"`
	// UnitPromotion         struct {
	// 	UnitID         int `json:"unitId"`
	// 	PromotionLevel int `json:"promotionLevel"`
	// 	EquipSlot1     int `json:"equipSlot_1"`
	// 	EquipSlot2     int `json:"equipSlot_2"`
	// 	EquipSlot3     int `json:"equipSlot_3"`
	// 	EquipSlot4     int `json:"equipSlot_4"`
	// 	EquipSlot5     int `json:"equipSlot_5"`
	// 	EquipSlot6     int `json:"equipSlot_6"`
	// }
	IsFinished      bool `json:"is_finished"`
	Position        int  `json:"position"`
	SearchAreaWidth int  `json:"search_area_width"`
}

func (Box) TableName() string {
	return "boxs"
}
