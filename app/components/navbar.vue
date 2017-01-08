<template>

  <div class="navbar">

    <a href="/">Schedule Buddy</a>

    <ul>

      <li><router-link to="/courses">Courses</router-link></li>

      <li><router-link to="/schedules">Schedules</router-link></li>

    </ul>

  </div>

</template>

<script>

  export default {

    name: 'navbar'

  }

</script>

<style lang="postcss" scoped>

  .navbar {

    background: #333;
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
