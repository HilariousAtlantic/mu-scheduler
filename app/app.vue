<template>

  <div class="app">

    <navbar></navbar>

    <router-view></router-view>

  </div>

</template>

<script>

  import Navbar from './components/navbar.vue';

  export default {

    name: 'app',

    components: {Navbar}

  }

</script>

<style lang="postcss" scoped>

  .app {

    height: 100vh;
    display: flex;
    flex-direction: column;

  }

</style>
