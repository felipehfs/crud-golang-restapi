<template>
  <div id="app">
    <Menu />
    <router-view/>
  </div>
</template>

<script>
import Menu from './components/Menu'

export default {
  name: 'App',
  components: {Menu}
}
</script>

<style>

  #app {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
</style>
