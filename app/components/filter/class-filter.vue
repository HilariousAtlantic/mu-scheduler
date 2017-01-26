<template>

    <schedule-filter :id="filter.id" :text="text" :active="filter.active">

      <option-input
        :options="operators"
        :defaultOption="{text: filter.options.operator}"
        @select="handleOperatorChange"
      ></option-input>

      <number-input
        :step="1"
        :defaultNumber="filter.options.amount"
        @change="handleAmountChange"
      ></number-input>

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
  import NumberInput from '../input/number-input.vue';
  import DaysInput from '../input/days-input.vue';

  import {formatDayList} from '../../lib/days';

  export default {

    name: 'class-filter',

    props: ['filter'],

    components: {ScheduleFilter, OptionInput, NumberInput, DaysInput},

    data() {

      return {

        operators: [

          {text: 'At Least'},

          {text: 'At Most'},

          {text: 'Exactly'}

        ]


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

        this.submitChanges({operator: operator.text});

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

        this.$store.dispatch('changeFilter', {id, changes});

      }

    }

  }

</script>

<style scoped>

  .option-input {

    width: 125px;

  }

</style>
