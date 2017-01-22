<template>

    <schedule-filter :id="filter.id" :text="text" :active="filter.active">

      <dropdown
        :options="['Start Before', 'Start After', 'Finish Before', 'Finish After']"
        :defaultOption="filter.options.operator"
        @select="handleOperatorChange"
      ></dropdown>

      <time-input
        :step="15"
        :defaultTime="filter.options.time"
        @change="handleTimeChange"
      ></time-input>

      <days-input
        :days="['M', 'T', 'W', 'R', 'F']"
        :defaultDays="filter.options.days"
        @change="handleDaysChange"
      ></days-input>

    </schedule-filter>

</template>

<script>

  import ScheduleFilter from './schedule-filter.vue';
  import Dropdown from '../common/dropdown.vue';
  import TimeInput from '../common/time-input.vue';
  import DaysInput from '../common/days-input.vue';

  import {formatDayList} from '../../lib/days';
  import {toTime} from '../../lib/time';

  export default {

    name: 'time-filter',

    props: ['filter'],

    components: {ScheduleFilter, Dropdown, TimeInput, DaysInput},

    data() {

      return {


      }

    },

    computed: {

      text() {

        let operator = this.filter.options.operator.toLowerCase();

        let time = toTime(this.filter.options.time);

        let days = formatDayList(this.filter.options.days);

        return 'I want ' + operator + ' ' + time + ' on ' + days;

      }

    },

    methods: {

      handleOperatorChange(operator) {

        this.submitChanges({operator});

      },

      handleTimeChange(time) {

        this.submitChanges({time});

      },


      handleDaysChange(days) {

        this.submitChanges({days});

      },

      submitChanges(change) {

        let id = this.filter.id;
        let changes = Object.assign({}, this.filter.options, change);

        this.$store.dispatch('changeScheduleFilter', {id, changes});

      }

    }

  }

</script>

<style scoped>

  .time-filter {

  }

</style>
