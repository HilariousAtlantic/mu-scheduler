<template>

  <div class="schedule-list">

    <double-header>

      <div slot="left">

        <button class="schedule-more" @click="toggleDetails">More Details</button>

      </div>

      <div slot="right">

        <button @click="$emit('prev')">Prev</button>

        <button @click="$emit('next')">Next</button>

      </div>

    </double-header>

    <schedule-details v-if="showDetails"
      :courses="schedule.courses"
      :gpa="schedule.gpa"
    ></schedule-details>

    <schedule-calendar
      :courses="schedule.courses"
      :start="schedule.start"
      :length="schedule.length"
    ></schedule-calendar>

  </div>

</template>

<script>

  import DoubleHeader from '../common/double-header.vue';
  import ScheduleDetails from './schedule-details.vue';
  import ScheduleCalendar from './schedule-calendar.vue';

  export default {

    name: 'schedule-list',

    props: ['index'],

    components: {DoubleHeader, ScheduleDetails, ScheduleCalendar},

    data() {

      return {

        showDetails: false

      }

    },

    computed: {

      schedule() {

        return this.$store.getters.filteredSchedules[this.index];

      }

    },

    methods: {

      toggleDetails() {

        this.showDetails = !this.showDetails;

      }

    }

  }

</script>

<style scoped>

  .double-header {

    margin-bottom: 10px;

  }

  button {

    border: 1px solid #ddd;
    padding: 10px 50px;
    background: #eee;
    outline: none;
    cursor: pointer;

    &:hover {

      background: #ddd;

    }

  }

  .schedule-more {

    background: #1565C0;
    color: #fff;

    &:hover {

      background: #0D47A1;

    }

  }

  .schedule-details {

    margin-bottom: 10px;

  }

</style>
