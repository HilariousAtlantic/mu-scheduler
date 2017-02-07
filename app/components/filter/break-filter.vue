<template>

    <schedule-filter :id="filter.id" :text="text" :active="filter.active">

      <time-input
        :step="15"
        :defaultTime="filter.options.start"
        @change="handleStartChange"
      ></time-input>

      <time-input
        :step="15"
        :defaultTime="filter.options.finish"
        @change="handleFinishChange"
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
  import TimeInput from '../input/time-input.vue';
  import DaysInput from '../input/days-input.vue';

  import {formatDayList} from '../../lib/days';
  import {toTime} from '../../lib/time';

  export default {

    name: 'break-filter',

    props: ['filter'],

    components: {ScheduleFilter, TimeInput, DaysInput},

    computed: {

      text() {

        let start = toTime(this.filter.options.start);

        let finish = toTime(this.filter.options.finish);

        let days = formatDayList(this.filter.options.days);

        return 'I want a break from ' + start + ' until ' + finish + ' on ' + days;

      }

    },

    methods: {

      handleStartChange(start) {

        this.submitChanges({start});

      },

      handleFinishChange(finish) {

        this.submitChanges({finish});

      },

      handleDaysChange(days) {

        this.submitChanges({days});

      },

      submitChanges(change) {

        let id = this.filter.id;
        let changes = Object.assign({}, this.filter.options, change);

        this.$store.dispatch('changeFilter', {id, changes});

      }

    }

  }

</script>

<style scoped>


</style>
