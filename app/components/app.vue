<template>

  <div class="app">

    <navbar></navbar>

    <div class="container">

      <generator></generator>

      <router-view></router-view>

    </div>

  </div>

</template>

<script>

  import Navbar from './common/navbar.vue';
  import Generator from './generator.vue';

  export default {

    name: 'app',

    components: {Navbar, Generator},

    mounted() {

      this.$store.dispatch('fetchTerms');

    }

  }

</script>

<style>

  .app {

    color: #333;
    margin-bottom: 50px;

  }

  .container {

    width: 90%;
    max-width: 1000px;
    margin: 0 auto;

  }

</style>
