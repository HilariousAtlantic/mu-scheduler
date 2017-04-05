<template>

  <div class="generator" @click="$store.dispatch('selectEditor', '')">

    <div class="navbar">

      <a href="/">Miami Scheduler</a>

    </div>

    <div class="view">

      <router-view></router-view>

    </div>

  </div>

</template>

<script>

  export default {

    name: 'generator',

    mounted() {

      this.$store.dispatch('fetchTerms');

    }

  }

</script>

<style>

  .navbar {

    background: #B61E2E;
    padding: 15px;

    a {
      font-size: 1.25rem;
      text-decoration: none;
      color: #f5f5f5;
    }

  }

  .view {
    width: 90%;
    max-width: 950px;
    margin: 20px auto;
  }

</style>
