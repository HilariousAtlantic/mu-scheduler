<template>

  <div class="finish-filter">

    <input
      type="checkbox"
      class="form-checkbox"
    >

    <option-input
      :options="operators"
      :defaultValue="changes.operator"
      :renderOption="option => option"
      @change="handleOperatorChange"
    ></option-input>

    <time-input
      :defaultValue="changes.time"
      @change="handleTimeChange"
    ></time-input>

    <days-input
      :defaultValue="changes.days"
      @change="handleDaysChange"
    ></days-input>

  </div>

</template>

<script>

  import OptionInput from '../common/option-input.vue';
  import TimeInput from '../common/time-input.vue';
  import DaysInput from '../common/days-input.vue';

  export default {

    name: 'finish-filter',

    props: ['active', 'options'],

    components: {OptionInput, TimeInput, DaysInput},

    data() {

      return {

        operators: ['Before', 'Exactly', 'After'],

        changes: {...this.options}

      }

    },

    methods: {

      handleOperatorChange(operator) {

        this.changes.operator = operator;
        this.$emit('change', this.changes);

      },

      handleTimeChange(time) {

        this.changes.time = time;
        this.$emit('change', this.changes);

      },


      handleDaysChange(days) {

        this.changes.days = days;
        this.$emit('change', this.changes);

      }

    }

  }

</script>

<style scoped>

  .finish-filter {
    display: flex;
    align-items: center;
  }

  .finish-filter > * + * {
    margin-left: 10px;
  }

</style>
