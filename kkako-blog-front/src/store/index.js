import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    user: {},
    width: 0,
  },
  mutations: {
    login(state, user) {
      state.user = user
    },
    logout(state) {
      state.user = {}
    },
    setWidth(state, width) {
      state.width = width
    }
  },
  actions: {
  },
  modules: {
  }
})
