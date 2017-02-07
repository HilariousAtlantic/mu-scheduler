<template>

  <div class="app">

    <navbar></navbar>

    <router-view></router-view>

  </div>

</template>

<script>

  import Navbar from './navbar.vue';

  export default {

    name: 'app',

    components: {Navbar},

    mounted() {

      this.$store.dispatch('fetchTerms');
      this.$store.dispatch('createFilter', 'time');
      this.$store.dispatch('createFilter', 'class');
      this.$store.dispatch('createFilter', 'break');

    }

  }

</script>

<style>

  .app {

    color: #333;
    margin-bottom: 50px;

  }

</style>
