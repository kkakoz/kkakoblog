<template>
  <div id="app">
    <nav-bar></nav-bar>
    <div id="main-container">
      <div id="main">
        <!-- <div>{{$store.state.width}}</div> -->
        <router-view/>
      </div>
    </div>
    <footers></footers>
    
  </div>
</template>

<script>
import Footers from './components/Footers.vue';

import HeaderMenu from './components/HeaderMenu.vue';
import NavBar from './views/NavBar.vue';

export default {
  name: 'app',
  components: {
    HeaderMenu,
    NavBar,
    Footers
  },
  data () {
  },
  methods: {
  },
  mounted() {
    this.$store.commit('setWidth', document.body.clientWidth)
    window.onresize = () => {
        this.$store.commit('setWidth', document.body.clientWidth)
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  height: 100%;
  background-color: #f9f9f9;
}

#main-container {
  width:100%;
  
}

#main {
  /* margin: 10px 20% 0 20%; */
  padding: 0;
  width: 80%;
  margin: 10px auto;
}

#main-container a {
  font-weight: bold;
  color: #2c3e50;
}

#main-container a.router-link-exact-active {
  color: #42b983;
}

.close {
    position: absolute;
    right: 0;
    top: 0;
}
</style>
