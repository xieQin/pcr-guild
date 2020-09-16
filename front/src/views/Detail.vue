<template>
<div>
  <img v-for="img in imgs" :key="img" class="img" v-lazy="img">
</div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { Lazyload } from 'vant'

@Component
export default class Home extends Vue {
  imgs = [
    "http://pcr.pekopeko.world/img/details-1.jpeg",
    "http://pcr.pekopeko.world/img/details-2.jpeg",
    "http://pcr.pekopeko.world/img/details-3.jpeg"
  ]
}
</script>