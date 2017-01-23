<template>

  <div class="generator">

    <div class="tab-list">

      <router-link class="tab" to="/courses">

        <span>Courses</span>

      </router-link>

      <router-link class="tab" to="/filters">

        <span>Filters</span>

      </router-link>

      <router-link class="tab" to="/schedules">

        <span>Schedules</span>

      </router-link>

    </div>

  </div>

</template>

<script>

  export default {

    name: 'generator'

  }

</script>

<style>

  .generator {

    width: 90%;
    max-width: 1000px;
    margin: 10px auto;

  }

  .tab-list {

    display: flex;
    align-items: center;
    text-align: center;

  }

  .tab {

    flex: 1;
    padding: 10px;
    border-bottom: 1px solid #B61E2E;
    text-decoration: none;
    color: inherit;
    cursor: pointer;

    &:hover {

      background: #f5f5f5;

    }

  }

  .router-link-active {

    border-bottom: 4px solid #B61E2E;

  }

</style>
