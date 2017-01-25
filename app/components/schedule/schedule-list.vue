<template>

  <div class="schedule-list">

    <double-header>

      <div slot="left">

        <button class="schedule-select">Save Schedule</button>

      </div>

      <div slot="right">

        <button @click="$emit('prev')">Prev</button>

        <button @click="$emit('next')">Next</button>

      </div>

    </double-header>

    <schedule-calendar
      :courses="schedule.courses"
      :start="schedule.start"
      :length="schedule.length"
    ></schedule-calendar>

  </div>

</template>

<script>

  import DoubleHeader from '../common/double-header.vue';
  import ScheduleCalendar from './schedule-calendar.vue';

  export default {

    name: 'schedule-list',

    props: ['index'],

    components: {DoubleHeader, ScheduleCalendar},

    computed: {

      schedule() {

        return this.$store.getters.filteredSchedules[this.index];

      }

    }

  }

</script>

<style scoped>

  .schedule-list {

    display: flex;
    flex-direction: column;

  }

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

  .schedule-select {

    background: #4CAF50;
    color: #fff;

    &:hover {

      background: #43A047;

    }

  }

  .schedule {

    flex: 1;

  }

</style>
