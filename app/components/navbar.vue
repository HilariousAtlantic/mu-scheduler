<template>

  <div class="navbar">

    <a href="/">Schedule Buddy</a>

    <ul>

      <li class="active">Courses</li>

      <li>Schedules</li>

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
    color: #f5f5f5;
    line-height: 1.25rem;
    margin-left: 10px;

  }

  li.active {

    border-bottom: 1px solid #f5f5f5;

  }

</style>
