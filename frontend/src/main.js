import Vue from 'vue'
import App from './App.vue'
import router from './router/router'
import store from './store'
import './plugins/element.js'
import './plugins/init.js'
import './styles/index.scss'
Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app');

// eslint-disable-next-line no-console
console.log("ApiHub open source project: https://github.com/haozzzzzzzz/api_hub");