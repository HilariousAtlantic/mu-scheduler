<template>

  <div class="time-filter">

    <schedule-filter :text="text" :active="active"></schedule-filter>

  </div>

</template>

<script>

  import ScheduleFilter from './schedule-filter.vue';

  export default {

    name: 'time-filter',

    props: ['options', 'active'],

    components: {ScheduleFilter},

    computed: {

      text() {

        return 'I want to ' + this.options.operator.toLowerCase() + ' ' + this.options.time + ' on ' + this.options.days;

      }

    }

  }

</script>

<style scoped>

  .time-filter {

  }

</style>
