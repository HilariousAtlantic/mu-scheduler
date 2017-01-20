<template>

  <div class="schedule-list">

    <schedule-browser :index="index" @change="updateIndex"></schedule-browser>

    <schedule :courses="currentSchedule.courses"></schedule>

  </div>

</template>

<script>

  import Schedule from './schedule.vue';
  import ScheduleBrowser from './schedule-browser.vue';

  export default {

    name: 'schedule-list',

    components: {Schedule, ScheduleBrowser},

    data() {

      return {

        index: 0

      }

    },

    computed: {

      currentSchedule() {

        return this.$store.getters.filteredSchedules[this.index];

      }

    },

    methods: {

      updateIndex(index) {

        this.index = index;

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
