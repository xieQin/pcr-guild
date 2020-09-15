<template>
<div class="me">
  <van-cell-group style="margin-top: 20px;">
    <van-cell title="Box" is-link />
    <van-cell title="阵容" is-link />
    <van-cell title="角色" is-link />
  </van-cell-group>
</div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { Cell, CellGroup } from 'vant'

@Component({
  components: { Cell, CellGroup }
})
export default class Home extends Vue {
}
</script>