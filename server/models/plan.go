package models

type Plan struct {
	ID         int           `json:"id" gorm:"primary_key"`
	BattleID   int           `json:"battle_id"`
	Cycle      int           `json:"cycle"`
	BossNum    int           `json:"boss_num"`
	BossStage  int           `json:"boss_stage"`
	Damage     int           `json:"damage"`
	Day        int           `json:"day"`
	UserID     int           `json:"user_id"`
	TeamID     int           `json:"team_id"`
	TeamInfo   Team          `json:"team_info" gorm:"ForeignKey:TeamID"`
	BossID     int           `json:"boss_id"`
	BossInfo   Boss          `json:"boss_info" gorm:"ForeignKey:BossID"`
	AssistID   int           `json:"assist_id"`
	AssistInfo TeamCharacter `json:"assist_info" gorm:"ForeignKey:AssistID"`
	IsEnd      bool          `json:"is_end"`      // 是否尾刀
	IsContinue bool          `json:"is_continue"` // 是否补偿刀
}

func (Plan) TableName() string {
	return "plans"
}
