<template>

  <div v-else class="schedule-list">

    <schedule-calendar
      :courses="schedule.courses"
      :start="schedule.start"
      :length="schedule.length"
    ></schedule-calendar>

  </div>

</template>

<script>

  import ScheduleCalendar from './schedule-calendar.vue';

  export default {

    name: 'schedule-list',

    props: ['index', 'showDetails'],

    components: {ScheduleCalendar},

    computed: {

      schedule() {

        return this.$store.getters.filteredSchedules[this.index];

      }

    }

  }

</script>

<style scoped>

</style>
