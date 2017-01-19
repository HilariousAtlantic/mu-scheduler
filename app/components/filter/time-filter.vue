<template>

  <div class="time-filter">

    <schedule-filter :text="text" :active="active"></schedule-filter>

  </div>

</template>

<script>

  import ScheduleFilter from './schedule-filter.vue';

  export default {

    name: 'time-filter',

    props: ['operator', 'time', 'days', 'active'],

    components: {ScheduleFilter},

    computed: {

      text() {

        return 'I want to ' + this.operator.toLowerCase() + ' ' + this.time + ' on ' + this.days;

      }

    }

  }

</script>

<style scoped>

  .time-filter {

  }

</style>
