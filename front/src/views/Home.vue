<template>
  <div class="content home">
    <h2 class="title">角色</h2>
    <icon name="home-o"/>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { Icon } from 'vant'

@Component({
  components: { Icon }
})
export default class Home extends Vue {
}
</script>
