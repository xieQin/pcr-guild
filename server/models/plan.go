package models

type Plan struct {
	ID       int `json:"id"`
	Day      int `json:"day"`
	TeamInfo struct {
		Team []PlanTeam `json:"plan_team"`
	} `json:"team_info"`
	BossInfo Boss `json:"boss_info"`
}

type PlanTeam struct {
	Team         Team
	TargetDamage int `json:"target_damage"`
	Num          int `json:"num"`
}
