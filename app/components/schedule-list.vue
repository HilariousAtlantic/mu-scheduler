<template>

  <div class="schedule-list">

    <div class="schedule-toolbar">

      <div class="schedule-search">

        <button @click="decreaseIndex"><i class="fa fa-angle-left"></i></button>
        <span>{{index+1}} of {{schedules.length}}</span>
        <button @click="increaseIndex"><i class="fa fa-angle-right"></i></button>

      </div>

      <button class="schedule-select">Save Schedule</button>

    </div>

    <div class="schedule-wrapper">

      <span v-if="noSchedules" class="no-schedules">No schedules generated</span>

      <div v-else-if="loadingSchedules" class="loader"></div>

      <schedule v-else :courses="currentSchedule.courses"></schedule>

    </div>

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

  .schedule-wrapper {

    flex: 1;
    display: flex;
    justify-content: center;

  }

  .schedule {

    flex: 1;

  }

  .no-schedules {

    flex: 1;
    align-self: center;
    text-align: center;
    font-size: 2rem;
    color: #999;

  }

  .loader {

    align-self: center;
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
