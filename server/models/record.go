package models

type Record struct {
	ID       int  `json:"id"`
	Day      int  `json:"day"`
	TeamInfo Team `json:"team_info"`
	BossInfo Boss `json:"boss_info"`
	UserID   int  `json:"user_id"`
	Damage   int  `json:"damage"`
}
