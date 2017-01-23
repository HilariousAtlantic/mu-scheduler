<template>

  <div class="navbar">

    <a href="/">Course Scheduler for Miami University</a>

    <ul>

      <li><router-link to="/">Generator</router-link></li>

    </ul>

  </div>

</template>

<script>

  export default {

    name: 'navbar'

  }

</script>

<style scoped>

  .navbar {

    background: #B61E2E;
    padding: 20px;

  }

  .router-link-active {

    border-bottom: 1px solid #f5f5f5;

  }

  a {

    font-size: 1.25rem;
    text-decoration: none;
    color: #f5f5f5;

  }

  ul {

    float: right;
    margin: 0;
    padding: 0;

  }

  li {

    display: inline-block;
    list-style: none;
    margin-left: 10px;

  }

  li a {

    font-size: 1rem;
    line-height: 1.25rem;

  }

</style>
