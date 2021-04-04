import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import axios from 'axios'
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
import 'element-ui/lib/theme-chalk/display.css';

Vue.use(mavonEditor)
Vue.config.productionTip = false
Vue.use(ElementUI);

axios.defaults.baseURL = "localhost:10003"
axios.defaults.timeout = 5000 // ms

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
