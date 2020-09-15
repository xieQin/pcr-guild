import { ICharacter } from "./character"

export interface ITeam {
  team_id: number
  boss_num: number
  boss_stage: number
  team_name: string
  target_damage: number
  context: string
  character_info: {
    tid: number
    level: number
    rarity: number
    promotion: number
    love_level: number
    slot1: boolean
    slot2: boolean
    slot3: boolean
    slot4: boolean
    slot5: boolean
    slot6: boolean
    unique_equip_rank: number
    team_id: number
    character_id: number
    character: ICharacter
  }
}
