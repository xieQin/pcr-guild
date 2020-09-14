import Vue from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'
import store from './store'
import api from './apis/axios'
import 'vant/lib/button/style'
import './style.less'

import { Icon, TabbarItem } from 'vant'

Vue.use(Icon)
Vue.use(TabbarItem)

Vue.config.productionTip = false
Vue.prototype.$api = api

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
