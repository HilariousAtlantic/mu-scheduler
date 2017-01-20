<template>

  <div class="time-filter-options">

    <dropdown
      :options="['Start Before', 'Start After', 'Finish Before', 'Finish After']"
      :defaultOption="options.operator"
      @select="onOperatorChange"
    ></dropdown>

    <time-input
      :step="15"
      :defaultTime="options.time"
      @change="onTimeChange"
    ></time-input>

    <days-input
      :days="['M', 'T', 'W', 'R', 'F', 'S']"
      :defaultDays="options.days"
      @change="onDaysChange"
    ></days-input>

  </div>

</template>

<script>

  import Dropdown from '../common/dropdown.vue';
  import TimeInput from '../common/time-input.vue';
  import DaysInput from '../common/days-input.vue';

  export default {

    name: 'time-filter-options',

    props: ['options'],

    components: {Dropdown, TimeInput, DaysInput},

    methods: {

      onOperatorChange(operator) {

        this.$emit('change', {operator});

      },

      onTimeChange(time) {

        this.$emit('change', {time});

      },

      onDaysChange(days) {

        this.$emit('change', {days});

      }

    }

  }

</script>

<style scoped>

  .time-filter-options {

    display: inline-block;

  }

</style>
