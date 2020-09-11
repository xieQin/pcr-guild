package models

type Team struct {
	BossNum       int    `json:"boss_num"`
	BossStage     int    `json:"boss_stage"`
	TeamName      string `json:"team_name"`
	Damage        int    `json:"target_damage"`
	CharacterInfo struct {
		Character []TeamCharacter `json:"team_character"`
	} `json:"character_info"`
	AssistInfo Character `json:"assist_character"`
}

type TeamCharacter struct {
	Character       Character
	Level           int  `json:"level"`     // 等级
	Rarity          int  `json:"rarity"`    //星级
	Promotion       int  `json:"promotion"` // rank
	LoveLevel       int  `json:"loveLevel"` // 好感度
	Slot1           bool `json:"slot1"`
	Slot2           bool `json:"slot2"`
	Slot3           bool `json:"slot3"`
	Slot4           bool `json:"slot4"`
	Slot5           bool `json:"slot5"`
	Slot6           bool `json:"slot6"`
	UniqueEquipRank int  `json:"uniqueEquipRank"` // 专武
}
