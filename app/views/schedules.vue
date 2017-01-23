<template>

  <div class="schedules-view">

    <double-header>

      <span slot="left">Generated Schedules</span>

      <span slot="right">{{numSchedules}} Schedules</span>

    </double-header>

    <schedule-list></schedule-list>

    <router-link v-if="numSchedules === 0" to="/course">

      <button class="courses-select">Select Courses</button>

    </router-link>

    <router-link v-else to="/filters">

      <button class="filters-change">Change Filters</button>

    </router-link>

    </div>

  </div>

</template>

<script>

  import DoubleHeader from '../components/common/double-header.vue';
  import ScheduleList from '../components/schedule/schedule-list.vue';

  export default {

    name: 'schedules-view',

    components: {DoubleHeader, ScheduleList},

    computed: {

      numSchedules() {

        return this.$store.state.schedules.length;

      }

    }

  }

</script>

<style scoped>

  .schedules-view {

    width: 90%;
    max-width: 1000px;
    margin: 20px auto;

  }

  .double-header {

    font-weight: 900;
    font-size: 1.25rem;

  }

  .schedule-list {

    margin: 20px 0;

  }

  .courses-select, .filters-change {

    width: 100%;
    padding: 15px;
    border: 1px solid #ddd;
    background: #4CAF50;
    color: #fff;
    outline: none;
    cursor: pointer;

  }

</style>
