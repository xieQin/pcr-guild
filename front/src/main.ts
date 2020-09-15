import Vue from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'
import store from './store'
import api from './apis/axios'
import 'vant/lib/button/style'
import './style.less'

import { Col, Row, Icon, TabbarItem, Button, Form, Field, RadioGroup, Radio, Stepper, Step, Steps, Cell, CellGroup } from 'vant'

Vue.use(CellGroup)
Vue.use(Cell)
Vue.use(Icon)
Vue.use(TabbarItem)
Vue.use(Col)
Vue.use(Row)
Vue.use(Button)
Vue.use(Form)
Vue.use(Field)
Vue.use(RadioGroup)
Vue.use(Radio)
Vue.use(Stepper)
Vue.use(Step)
Vue.use(Steps)

Vue.config.productionTip = false
Vue.prototype.$api = api

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
