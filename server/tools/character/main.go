package main

import (
	m "pcr-guild/models"
	u "pcr-guild/utils"
)

func main() {
	boxData := u.ReadFile("../../data/data.json", 1024*1024)
	var (
		boxs []m.Box
	)
	u.Unmarshal(boxData, &boxs)

	res := make([]m.Character, 0)
	for _, box := range boxs {
		item := m.Character{
			ID:       box.ID,
			UnitID:   box.UnitID,
			UnitName: box.UnitName,
			Position: box.Position,
			Icon:     box.Icon,
		}
		res = append(res, item)
	}
	path := "../../data/character.json"
	u.RemoveFile(path)
	u.WriteFile(path, u.Marshal(res))
}
