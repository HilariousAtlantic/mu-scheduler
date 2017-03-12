<template>

  <div class="schedules-view">

    <filter-list></filter-list>

    <schedule-toolbar :index="index" @dec="decreaseIndex" @inc="increaseIndex"></schedule-toolbar>

    <schedule-list :index="index"></schedule-list>

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

        index: 0

      }

    },

    methods: {

      decreaseIndex() {

        if (this.index > 0) {

          this.index--;

        }

      },

      increaseIndex() {

        if (this.index < this.$store.state.schedules.length-1) {

          this.index++;

        }

      }

    }

  }

</script>

<style scoped>

  .filter-list  {
    margin-bottom: 10px;
  }

  .schedule-toolbar {
    margin-bottom: 10px;
  }

  .schedule-list {
    margin-bottom: 10px;
  }

  .btn-primary {
    font-weight: normal;
  }

</style>
