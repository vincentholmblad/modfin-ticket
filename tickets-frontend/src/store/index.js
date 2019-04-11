import Vue from 'vue'
import Vuex from 'vuex'
import tickets from './modules/tickets'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
  modules: {
    tickets
  },
  strict: debug
})