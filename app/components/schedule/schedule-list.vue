<template>

  <div v-if="noSchedules" class="no-schedules">

    <h3>No schedules generated</h3>

    <router-link to="/courses"><button>Select Courses</button></router-link>

  </div>

  <div v-else-if="loadingSchedules" class="loader-wrapper">

    <div class="loader"></div>

  </div>

  <div v-else class="schedule-list">

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

        return this.schedules[0];

      },

      loadingSchedules() {

        return this.$store.state.requestingSchedules;

      },

      noSchedules() {

        return this.$store.state.schedules.length === 0;

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

  .no-schedules {

    flex: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;


    h3 {

      font-size: 2rem;
      color: #ccc;
      text-transform: uppercase;

    }

    button {

      width: 400px;
      padding: 10px;
      font-size: 1rem;
      border: 1px solid #ddd;
      background: #1565C0;
      color: #fff;
      outline: none;
      cursor: pointer;

      &:hover {

        background: #0D47A1;

      }

    }

  }

  .loader-wrapper {

    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;

  }

  .loader {

    width: 100px;
    height: 100px;
    border: 10px solid #ddd;
    border-bottom: 10px solid #444;
    border-radius: 90px;
    animation: spin 1s linear infinite;

  }

  @keyframes spin {

    to {

      transform: rotate(360deg);

    }

  }

</style>
