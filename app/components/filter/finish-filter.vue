<template>

  <div class="finish-filter filter">

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

    <button class="btn btn-danger" @click="handleFilterDelete">Remove</button>

  </div>

</template>

<script>

  import OptionInput from '../common/option-input.vue';
  import TimeInput from '../common/time-input.vue';
  import DaysInput from '../common/days-input.vue';

  export default {

    name: 'finish-filter',

    props: ['id', 'options'],

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
        this.submitChanges();

      },

      handleTimeChange(time) {

        this.changes.time = time;
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

  .finish-filter {
    display: flex;
    align-items: center;
  }

  .finish-filter > * + * {
    margin-left: 10px;
  }

</style>
