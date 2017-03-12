<template>

  <div class="class-filter filter">

    <option-input
      :options="operators"
      :defaultValue="changes.operator"
      :renderOption="option => option"
      @change="handleOperatorChange"
    ></option-input>

    <number-input
      :defaultValue="changes.amount"
      @change="handleAmountChange"
    ></number-input>

    <days-input
      :defaultValue="changes.days"
      @change="handleDaysChange"
    ></days-input>

    <button class="btn btn-danger" @click="handleFilterDelete">Remove</button>

  </div>

</template>

<script>

  import OptionInput from '../common/option-input.vue';
  import NumberInput from '../common/number-input.vue';
  import DaysInput from '../common/days-input.vue';

  export default {

    name: 'class-filter',

    props: ['id', 'options'],

    components: {OptionInput, NumberInput, DaysInput},

    data() {

      return {

        operators: ['At Least', 'Exactly', 'At Most'],

        changes: {...this.options}

      }

    },

    methods: {

      handleOperatorChange(operator) {

        this.changes.operator = operator;
        this.submitChanges();

      },

      handleAmountChange(amount) {

        this.changes.amount = amount;
        this.submitChanges();

      },


      handleDaysChange(days) {

        this.changes.days = days;
        this.submitChanges();

      },

      submitChanges() {

        let changes = {...this.changes};
        let id = this.id;

        this.$store.dispatch('changeFilter', {id, changes});

      },

      handleFilterDelete() {

        this.$store.dispatch('deleteFilter', this.id);

      }

    }

  }

</script>

<style scoped>

  .class-filter {
    display: flex;
    align-items: center;
  }

  .class-filter > * + * {
    margin-left: 10px;
  }

</style>
