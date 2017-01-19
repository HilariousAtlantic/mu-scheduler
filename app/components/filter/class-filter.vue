<template>

  <div class="class-filter">

    <schedule-filter :text="text" :active="options.active"></schedule-filter>

  </div>

</template>

<script>

  import ScheduleFilter from './schedule-filter.vue';

  export default {

    name: 'class-filter',

    props: ['options'],

    components: {ScheduleFilter},

    computed: {

      text() {

        return 'I want ' + this.options.operator.toLowerCase() + ' ' + this.options.amount + ' classes on ' + this.options.days;

      }

    }

  }

</script>

<style scoped>

  .class-filter {

  }

</style>
