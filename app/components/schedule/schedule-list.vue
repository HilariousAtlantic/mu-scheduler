<template>

  <div class="schedule-list">

    <schedule-browser></schedule-browser>

    <schedule :courses="currentSchedule.courses"></schedule>

  </div>

</template>

<script>

  import Schedule from './schedule.vue';
  import ScheduleBrowser from './schedule-browser.vue';

  export default {

    name: 'schedule-list',

    components: {Schedule, ScheduleBrowser},

    computed: {

      schedules() {

        return this.$store.getters.generatedSchedules;

      },

      currentSchedule() {

        return this.schedules[this.$store.state.schedulesIndex];

      }

    }

  }

</script>

<style scoped>

  .schedule-list {

    display: flex;
    flex-direction: column;

  }

  .schedule-browser {

    margin-bottom: 10px;

  }

  .schedule {

    flex: 1;

  }

</style>
