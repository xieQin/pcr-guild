<template>
<div>
  <img v-for="img in imgs" :key="img" class="img" v-lazy="img">
</div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { Lazyload } from 'vant'
import * as img1 from '@/assets/details-1.jpeg'
import * as img2 from  '@/assets/details-2.jpeg'
import * as img3 from '@/assets/details-3.jpeg'

@Component
export default class Home extends Vue {
  imgs = [
    img1,
    img2,
    img3
  ]
}
</script>