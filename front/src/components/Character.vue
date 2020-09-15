<template>
  <div class="character-box">
    <p class="tc name">{{unit_name}}</p>
    <p class="tc position">{{positionTM[position] || ""}}</p>
    <img class="m-auto" width="45" height="45" :src="`https://pcredivewiki.tw/static/images/unit/icon_unit_${icon}.png`" alt="">
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'

@Component
export default class Character extends Vue{
  @Prop({ type: String }) unit_name !: string
  @Prop({ type: Number, default: 0 }) position !: number
  @Prop({ type: String }) icon !: string

  positionTM = ["前卫", "中卫", "后卫"]
}
</script>

<style lang="less" scoped>
.character-box {
  padding: 8px;
}
.name {
  font-size: 13px;
  font-weight: bold;
  margin: 0;
  height: 15px;
  overflow: hidden;
}
.position {
  font-size: 12px;
  margin: 0;
  color: red;
}
</style>