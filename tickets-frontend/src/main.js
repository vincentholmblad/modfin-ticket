import Vue from 'vue'
import App from './App.vue'
import router from './router'

import axios from 'axios'
import VueAxios from 'vue-axios'

Vue.use(VueAxios, axios)

Vue.config.productionTip = false

import VueBus from 'vue-bus';
 
Vue.use(VueBus);

import { Drag, Drop } from 'vue-drag-drop';
 
Vue.component('drag', Drag);
Vue.component('drop', Drop);

import Vue2Filters from 'vue2-filters'

Vue.use(Vue2Filters)

Vue.use(require('vue-moment'));

require('@/assets/scss/main.scss');

var VueCookie = require('vue-cookie');
// Tell Vue to use the plugin
Vue.use(VueCookie);

new Vue({
  router,
  render: h => h(App),
  data: {
    userName: ''
  }
}).$mount('#app')
