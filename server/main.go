package main

import (
	"fmt"

	m "pcr-guild/models"
)

func main() {
	var (
		teams = make([]m.Team, 0)
	)

	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		teams = append(teams, m.Team{
			BossNum:   1,
			BossStage: 1,
			Damage:    i * 50,
		})
	}
	fmt.Print(teams)
}
