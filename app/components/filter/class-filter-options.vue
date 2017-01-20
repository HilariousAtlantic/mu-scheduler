<template>

  <div class="class-filter-options">

    <dropdown
      :options="['At Least', 'At Most', 'Exactly']"
      :defaultOption="options.operator"
      @select="onOperatorChange"
    ></dropdown>

    <number-input
      :step="1"
      :defaultNumber="options.amount"
      @change="onAmountChange"
    ></number-input>

    <days-input
      :days="['M', 'T', 'W', 'R', 'F', 'S']"
      :defaultDays="options.days"
      @change="onDaysChange"
    ></days-input>

  </div>

</template>

<script>

  import Dropdown from '../editor/dropdown.vue';
  import NumberInput from '../editor/number-input.vue';
  import DaysInput from '../editor/days-input.vue';

  export default {

    name: 'class-filter-options',

    props: ['options'],

    components: {Dropdown, NumberInput, DaysInput},

    methods: {

      onOperatorChange(operator) {

        this.$emit('change', {operator});

      },

      onAmountChange(amount) {

        this.$emit('change', {amount});

      },

      onDaysChange(days) {

        this.$emit('change', {days});

      }

    }

  }

</script>

<style scoped>

  .class-filter-options {

    display: inline-block;

  }

</style>
