import Vue from 'vue'
import router from './router'
import store from './store'
import axios from 'axios'
Vue.prototype.$axios = axios;
Vue.config.productionTip = false

import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import App from './App.vue';

Vue.use(ElementUI);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
