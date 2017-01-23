<template>

  <div class="app">

    <navbar></navbar>

    <generator></generator>

    <router-view class="page"></router-view>

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

    display: flex;
    flex-direction: column;
    color: #333;

  }

  .page {

    flex: 1;

  }

</style>
