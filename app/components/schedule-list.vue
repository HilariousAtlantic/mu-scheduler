<template>

  <div v-if="noSchedules" class="no-schedules">

    <h3>No schedules generated</h3>

    <router-link to="/courses"><button>Select Courses</button></router-link>

  </div>

  <div v-else-if="loadingSchedules" class="loader-wrapper">

    <div class="loader"></div>

  </div>



  <div v-else class="schedule-list">

    <div class="schedule-toolbar">

      <div class="schedule-search">

        <button @click="decreaseIndex"><i class="fa fa-angle-left"></i></button>
        <span>{{index+1}} of {{schedules.length}}</span>
        <button @click="increaseIndex"><i class="fa fa-angle-right"></i></button>

      </div>

      <button class="schedule-select">Save Schedule</button>

    </div>

    <schedule :courses="currentSchedule.courses"></schedule>

  </div>

</template>

<script>

  import Schedule from './schedule.vue';

  export default {

    name: 'schedule-list',

    components: {Schedule},

    data() {

      return {

        index: 0

      }

    },

    computed: {

      schedules() {

        return this.$store.getters.generatedSchedules;

      },

      currentSchedule() {

        return this.schedules[this.index];

      },

      loadingSchedules() {

        return this.$store.state.requestingSchedules;

      },

      noSchedules() {

        return this.$store.state.schedules.length === 0;

      }

    },

    methods: {

      increaseIndex() {

        if (this.index < this.schedules.length-1) {

          this.index++;

        }

      },

      decreaseIndex() {

        if (this.index > 0) {

          this.index--;

        }

      }

    }

  }

</script>

<style scoped>

  .schedule-list {

    display: flex;
    flex-direction: column;

  }

  .schedule-toolbar {

    margin-bottom: 10px;

  }

  .schedule-search {

    display: inline-block;

    button {

      border: 1px solid #ddd;
      background: #eee;
      padding: 10px 50px;
      outline: none;
      cursor: pointer;

    }

    span {

      text-align: center;
      border: 1px solid #ddd;
      padding: 10px 75px;

    }

  }

  .schedule-select {

    float: right;
    display: inline-block;
    padding: 10px 100px;
    border: 1px solid #ddd;
    background: #4CAF50;
    color: #fff;
    outline: none;
    cursor: pointer;

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
