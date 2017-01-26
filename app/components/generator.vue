<template>

  <div class="generator">

    <div class="tab-list">

      <router-link class="tab" to="courses">

        <span>Courses</span>

      </router-link>

      <router-link class="tab" to="filters">

        <span>Filters</span>

      </router-link>

      <router-link class="tab" to="schedules">

        <span>Schedules</span>

      </router-link>

    </div>

    <div class="view">

      <router-view></router-view>

    </div>

  </div>

</template>

<script>

  export default {

    name: 'generator'

  }

</script>

<style>

  .tab-list {

    display: flex;
    align-items: center;
    text-align: center;
    background: #A91121;

    .router-link-active {

      background: #B61E2E;

    }

  }

  .tab {

    flex: 1;
    padding: 15px;
    text-decoration: none;
    color: #fff;
    cursor: pointer;

    &:not(.router-link-active):hover {

      background: #9D0515;

    }

  }

  .view {

    width: 90%;
    max-width: 1200px;
    margin: 20px auto;

  }

</style>
