<template>

  <div class="app">

    <navbar></navbar>

    <router-view class="page"></router-view>

  </div>

</template>

<script>

  import Navbar from './common/navbar.vue';

  export default {

    name: 'app',

    components: {Navbar},

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
