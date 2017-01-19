<template>

  <div class="time-filter-editor">

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

  import Dropdown from '../editor/dropdown.vue';
  import TimeInput from '../editor/time-input.vue';
  import DaysInput from '../editor/days-input.vue';

  import {toTime} from '../../lib/time';

  export default {

    name: 'time-filter',

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

  .time-filter-editor {

    display: inline-block;

  }

</style>
