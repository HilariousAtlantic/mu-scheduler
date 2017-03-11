<template>

  <div class="generator">

    <div class="navbar">

      <a href="/">Course Scheduler for Miami University</a>

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
      this.$store.dispatch('createFilter', 'start');
      this.$store.dispatch('createFilter', 'finish');
      this.$store.dispatch('createFilter', 'break');
      this.$store.dispatch('createFilter', 'class');

    }

  }

</script>

<style>

  .navbar {

    background: #B61E2E;
    padding: 20px;

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
