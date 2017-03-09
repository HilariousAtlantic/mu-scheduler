<template>

  <div class="generator">

    <div class="tabnav">

      <button class="btn btn-sm btn-primary right">Generate Schedules</button>

      <nav class="tabnav-tabs">

        <router-link active-class="selected" class="tabnav-tab" to="courses">

          <span>Courses</span>

        </router-link>

        <router-link active-class="selected" class="tabnav-tab" to="filters">

          <span>Filters</span>

        </router-link>

        <router-link active-class="selected" class="tabnav-tab" to="schedules">

          <span>Schedules</span>

        </router-link>

      </nav>

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

  .generator {

    width: 90%;
    max-width: 800px;
    margin: 20px auto;

  }

  .btn-sm.btn-primary {

    font-weight: normal;
    background-image: linear-gradient(-180deg, #34d058 0%, #28a745 90%);

  }

</style>
