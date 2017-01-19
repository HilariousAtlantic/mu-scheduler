<template>

  <div class="class-filter">

    <schedule-filter :text="text" :active="active"></schedule-filter>

  </div>

</template>

<script>

  import ScheduleFilter from './schedule-filter.vue';

  export default {

    name: 'class-filter',

    props: ['operator', 'amount', 'days', 'active'],

    components: {ScheduleFilter},

    computed: {

      text() {

        return 'I want ' + this.operator.toLowerCase() + ' ' + this.amount + ' classes on ' + this.days;

      }

    }

  }

</script>

<style scoped>

  .class-filter {

  }

</style>
