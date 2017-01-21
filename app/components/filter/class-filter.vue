<template>

    <schedule-filter :id="filter.id" :text="text" :active="filter.active">

      <dropdown
        :options="['At Least', 'At Most', 'Exactly']"
        :defaultOption="filter.options.operator"
        @select="handleOperatorChange"
      ></dropdown>

      <number-input
        :step="1"
        :defaultNumber="filter.options.amount"
        @change="handleAmountChange"
      ></number-input>

      <days-input
        :days="['M', 'T', 'W', 'R', 'F', 'S']"
        :defaultDays="filter.options.days"
        @change="handleDaysChange"
      ></days-input>

    </schedule-filter>

</template>

<script>

  import ScheduleFilter from './schedule-filter.vue';
  import Dropdown from '../common/dropdown.vue';
  import NumberInput from '../common/number-input.vue';
  import DaysInput from '../common/days-input.vue';

  import {formatDayList} from '../../lib/days';

  export default {

    name: 'class-filter',

    props: ['filter'],

    components: {ScheduleFilter, Dropdown, NumberInput, DaysInput},

    data() {

      return {


      }

    },

    computed: {

      text() {

        let operator = this.filter.options.operator.toLowerCase();

        let amount = this.filter.options.amount;

        let days = formatDayList(this.filter.options.days);

        return 'I want ' + operator + ' ' + amount + ' classes on ' + days;

      }

    },

    methods: {

      handleOperatorChange(operator) {

        this.submitChanges({operator});

      },

      handleAmountChange(amount) {

        this.submitChanges({amount});

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

  .class-filter {

  }

</style>
