package models

type TeamCharacter struct {
	TID             int       `json:"tid" gorm:"primary_key"` // 阵容角色iD
	Level           int       `json:"level"`                  // 等级
	Rarity          int       `json:"rarity"`                 // 星级
	Promotion       int       `json:"promotion"`              // rank
	LoveLevel       int       `json:"love_level"`             // 好感度
	Slot1           bool      `json:"slot1"`
	Slot2           bool      `json:"slot2"`
	Slot3           bool      `json:"slot3"`
	Slot4           bool      `json:"slot4"`
	Slot5           bool      `json:"slot5"`
	Slot6           bool      `json:"slot6"`
	UniqueEquipRank int       `json:"unique_equip_rank"` // 专武等级
	TeamID          int       `json:"team_id"`           // 阵容ID
	CharacterID     int       `json:"character_id"`      // 角色ID
	CharacterInfo   Character `json:"character" gorm:"ForeignKey:CharacterID"`
}

type Team struct {
	TeamID       int             `json:"team_id" gorm:"primary_key"`
	BossNum      int             `json:"boss_num"`                                // boss
	BossStage    int             `json:"boss_stage"`                              // boss 阶段
	TeamName     string          `json:"team_name"`                               // 阵容名
	TargetDamage int             `json:"target_damage"`                           // 目标伤害
	Characters   []TeamCharacter `json:"character_info" gorm:"ForeignKey:TeamID"` // 阵容角色
	Context      string          `json:"context"`
}

func (TeamCharacter) TableName() string {
	return "team_characters"
}

func (Team) TableName() string {
	return "teams"
}
