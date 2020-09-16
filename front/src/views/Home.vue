<template>
<div>
  <van-swipe class="my-swipe" :autoplay="3000">
    <van-swipe-item v-for="(image, index) in images" @click="$router.push('detail')" :key="index">
      <img v-lazy="image" />
    </van-swipe-item>
  </van-swipe>
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
</div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import Character from '@/components/Character.vue'
import { MockCharacters } from '../mock/characters'
import { Row, Col, Swipe, SwipeItem, Lazyload } from 'vant'
import { ICharacter } from '../models/character'
import * as img1 from '@/assets/1.png'
import * as img2 from  '@/assets/2.png'

@Component({
  components: { Character, Row, Col, Swipe, SwipeItem }
})
export default class Home extends Vue {
  characters: ICharacter[] = MockCharacters
  front: ICharacter[] = []
  middle: ICharacter[] = []
  back: ICharacter[] = []
  images = [img1, img2]

  mounted() {
    this.front = this.characters.filter(i => i.position === 0)
    this.middle = this.characters.filter(i => i.position === 1)
    this.back = this.characters.filter(i => i.position === 2)
    // this.$api.get("character").then(res => {
      // this.characters = res.data.items
    // })
  }
}
</script>

<style>
  .my-swipe .van-swipe-item {
    color: #fff;
    font-size: 20px;
    line-height: 150px;
    text-align: center;
    background-color: #39a9ed;
  }
</style>