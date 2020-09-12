package models

// Box struct ...
type Box struct {
	ID                    int    `json:"id" gorm:"primary_key"`
	UserID                int    `json:"userId"`
	Server                string `json:"server"`
	UnitID                int    `json:"unitId"`
	UnitName              string `json:"unitName"`
	Trace                 bool   `json:"trace"`
	Level                 int    `json:"level"`
	Rarity                int    `json:"rarity"`
	Promotion             int    `json:"promotion"`
	LoveLevel             int    `json:"loveLevel"`
	Pieces                int    `json:"pieces"`
	Slot1                 bool   `json:"slot1"`
	Slot2                 bool   `json:"slot2"`
	Slot3                 bool   `json:"slot3"`
	Slot4                 bool   `json:"slot4"`
	Slot5                 bool   `json:"slot5"`
	Slot6                 bool   `json:"slot6"`
	UniqueEquipRank       int    `json:"uniqueEquipRank"`
	Icon                  string `json:"icon"`
	TargetRarity          int    `json:"targetRarity"`
	TargetPromotion       int    `json:"targetPromotion"`
	TargetLoveLevel       int    `json:"targetLoveLevel"`
	TargetUniqueEquipRank int    `json:"targetUniqueEquipRank"`
	UnitPromotion         struct {
		UnitID         int `json:"unitId"`
		PromotionLevel int `json:"promotionLevel"`
		EquipSlot1     int `json:"equipSlot_1"`
		EquipSlot2     int `json:"equipSlot_2"`
		EquipSlot3     int `json:"equipSlot_3"`
		EquipSlot4     int `json:"equipSlot_4"`
		EquipSlot5     int `json:"equipSlot_5"`
		EquipSlot6     int `json:"equipSlot_6"`
	}
	IsFinished      bool `json:"isFinished"`
	Position        int  `json:"position"`
	SearchAreaWidth int  `json:"searchAreaWidth"`
}

func (Box) TableName() string {
	return "boxs"
}
