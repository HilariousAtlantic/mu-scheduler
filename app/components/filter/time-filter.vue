<template>

    <schedule-filter :id="filter.id" :text="text" :active="filter.active">

      <option-input
        :options="operators"
        :defaultOption="{text: filter.options.operator}"
        @select="handleOperatorChange"
      ></option-input>

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
  import OptionInput from '../input/option-input.vue';
  import TimeInput from '../input/time-input.vue';
  import DaysInput from '../input/days-input.vue';

  import {formatDayList} from '../../lib/days';
  import {toTime} from '../../lib/time';

  export default {

    name: 'time-filter',

    props: ['filter'],

    components: {ScheduleFilter, OptionInput, TimeInput, DaysInput},

    data() {

      return {

        operators: [

          {text: 'Start Before'},

          {text: 'Start After'},

          {text: 'Finish Before'},

          {text: 'Finish After'},

        ]

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

        this.submitChanges({operator: operator.text});

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

  .option-input {

    width: 150px;

  }

</style>
