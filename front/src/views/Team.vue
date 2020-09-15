<template>
<div class="wrapper team">
  <van-button plain to="/addteam" type="default">添加阵容</van-button>
  <h2 class="title">阵容</h2>
  <div class="content">
    <div class="section">
      <Team
        :team_name="item.team_name" :target_damage="item.target_damage" :context="item.context"
        :boss_stage="item.boss_stage" :boss_num="item.boss_num"
        v-for="item in teams"
        :characters="item.character_info"
        :key="item.team_id"
      />
    </div>
  </div>
</div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import Team from '@/components/Team.vue'
import { Button } from 'vant'
import { ITeam } from '../models/team'
import { MockTeams } from '../mock/teams'

@Component({
  components: { Team, Button }
})
export default class Home extends Vue {
  teams = MockTeams
  mounted() {
    // this.$api.get("team").then(res => {
    //   this.teams = res.data.items
    // })
  }
}
</script>