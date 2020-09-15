<template>
  <div>
    <div class="character-box">
      <div class="tl">
        <span class="name">{{team_name}}</span>&nbsp;
        <span>{{boss_stage}}阶段</span>
        <span>{{boss_num}}号boss</span>
      </div>
      <p class="content">
        <span>目标伤害：{{target_damage}}w</span>
      </p>
      <p class="context">
        {{context}}
      </p>
      <van-row type="flex" justify="space-between">
        <van-col v-for="item in characters" :key="item.character.unit_name" span="8">
          <Character
            :unit_name="item.character.unit_name" :position="item.character.position" :icon="item.character.icon"
          />
        </van-col>
      </van-row>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import Character from './Character.vue'
import { Col, Row } from 'vant'

@Component({
  components: {
    Character, Col, Row
  }
})
export default class Team extends Vue{
  @Prop({ type: String }) team_name !: string
  @Prop({ type: Number, default: 0 }) target_damage !: number
  @Prop({ type: Number, default: 1 }) boss_num !: number
  @Prop({ type: Number, default: 1 }) boss_stage !: number
  @Prop({ type: String }) context !: string
  @Prop({ type: Array }) characters !: []
}
</script>

<style lang="less">
.character-box {
  padding: 8px;
}
.name {
  font-weight: bold;
  font-size: 18px;
}
</style>