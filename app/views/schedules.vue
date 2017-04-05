<template>

  <div class="schedules-view">

    <filter-list></filter-list>

    <div v-if="scheduleCount == 0" class="blankslate">

      <h2>No Schedules</h2>

      <p>You may have chosen courses that do not have any valid schedules</p>

    </div>

    <div v-else>

      <schedule-toolbar
        :index="index"
        :length="scheduleCount"
        @details="showDetails = !showDetails"
        @change="handleIndexChange"
      ></schedule-toolbar>

      <schedule-list :index="index"></schedule-list>

    </div>

    <router-link to="courses" class="btn btn-block btn-primary">

      <span>Select Courses</span>

    </router-link>

  </div>

</template>

<script>

  import FilterList from '../components/filter/filter-list.vue';
  import ScheduleToolbar from '../components/schedule/schedule-toolbar.vue';
  import ScheduleList from '../components/schedule/schedule-list.vue';

  export default {

    name: 'schedules-view',

    components: {FilterList, ScheduleToolbar, ScheduleList},

    data() {

      return {

        index: 0,

        showDetails: false

      }

    },

    computed: {

      scheduleCount() {

        return this.$store.getters.filteredSchedules.length;

      }

    },

    methods: {

      handleIndexChange(index) {

        if (index < 0) {

          this.index = 0;

        } else if (index > this.scheduleCount-1) {

          this.index = this.scheduleCount-1;

        } else {

          this.index = index;

        }

      }

    }

  }

</script>

<style scoped>

  .blankslate {
    padding: 50px 30px;
  }

  .filter-list, .schedule-toolbar, .schedule-list, .blankslate  {
    margin-bottom: 10px;
  }

  .btn-primary {
    padding: 10px;
  }

  @media screen and (max-width: 750px) {

    .filter-list {
      display: none !important;
    }

  }

</style>
