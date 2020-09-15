<template>
<div class="wrapper home">
  <h2 class="title">前卫</h2>
  <div class="content">
    <div class="section">
      <van-row type="flex" class="flex-wrap" gutter="10">
        <van-col v-for="item in front" :key="item.unit_name" span="6">
          <Character
            :unit_name="item.unit_name" :position="item.position" :icon="item.icon"
          />
        </van-col>
      </van-row>
    </div>
  </div>
  <h2 class="title">中卫</h2>
  <div class="content">
    <div class="section">
      <van-row type="flex" class="flex-wrap" gutter="10">
        <van-col v-for="item in middle" :key="item.unit_name" span="6">
          <Character
            :unit_name="item.unit_name" :position="item.position" :icon="item.icon"
          />
        </van-col>
      </van-row>
    </div>
  </div>
  <h2 class="title">后卫</h2>
  <div class="content">
    <div class="section">
      <van-row type="flex" class="flex-wrap" gutter="10">
        <van-col v-for="item in back" :key="item.unit_name" span="6">
          <Character
            :unit_name="item.unit_name" :position="item.position" :icon="item.icon"
          />
        </van-col>
      </van-row>
    </div>
  </div>
</div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import Character from '@/components/Character.vue'
import { MockCharacters } from '../mock/characters'
import { Row, Col } from 'vant'
import { ICharacter } from '../models/character'

@Component({
  components: { Character, Row, Col }
})
export default class Home extends Vue {
  characters: ICharacter[] = MockCharacters
  front: ICharacter[] = []
  middle: ICharacter[] = []
  back: ICharacter[] = []

  mounted() {
    this.front = this.characters.filter(i => i.position === 0)
    this.middle = this.characters.filter(i => i.position === 1)
    this.back = this.characters.filter(i => i.position === 2)
    console.log(this.front)
    // this.$api.get("character").then(res => {
      // this.characters = res.data.items
    // })
  }
}
</script>
