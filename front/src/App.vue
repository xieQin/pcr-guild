<template>
  <div id="app">
    <div id="nav">
      <nav-bar
        title="公主连结"
        @click-left="onClickLeft"
        @click-right="onClickRight"
      />
    </div>
    <div id="view">
      <router-view />
    </div>
    <div id="tab">
      <tabbar v-model="active">
        <tabbar-item replace to="/" icon="wap-home-o">首页</tabbar-item>
        <tabbar-item replace to="/team" icon="friends-o">阵容</tabbar-item>
        <tabbar-item replace to="/plan" icon="notes-o">排刀</tabbar-item>
        <tabbar-item replace to="/me" icon="setting-o">我的</tabbar-item>
      </tabbar>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { NavBar, Tabbar, TabbarItem, Icon } from 'vant'

@Component({
  components: { NavBar, Tabbar, TabbarItem, Icon }
})
export default class App extends Vue {
  active = 0
  onClickLeft() {}
  onClickRight() {}
}
</script>

<style lang="less" scoped>
#view {
  position: fixed;
  width: 100%;
  top: 46px;
  bottom: 46px;
  overflow: auto;
  -webkit-overflow-scrolling : touch;
  background-color: #f7f8fa;
}
#tab {
  position: fixed;
  bottom: 0;
  width: 100%;
}
</style>