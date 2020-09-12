package models

type Boss struct {
	BossID    int    `json:"boss_id" gorm:"primary_key"`
	BattleID  int    `json:"battle_id"`
	Battle    string `json:"battle"`
	BossName  string `json:"boss_name"`
	Cycle     int    `json:"cycle"`
	BossNum   int    `json:"boss_num"`
	BossStage int    `json:"boss_stage"`
	TotalHP   int    `json:"total_hp"`
	CurrentHP int    `json:"current_hp"`
	RemainHP  int    `json:"remain_hp"`
}

func (Boss) TableName() string {
	return "boss"
}
