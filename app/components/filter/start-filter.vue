<template>

  <div class="start-filter">

    <input
      type="checkbox"
      class="form-checkbox"
      @change="handleActiveToggle">

    <option-input
      :options="operators"
      :defaultValue="changes.operator"
      :renderOption="option => option"
      @change="handleOperatorChange">
    </option-input>

    <time-input
      :defaultValue="changes.time"
      @change="handleTimeChange">
    </time-input>

    <days-input
      :defaultValue="changes.days"
      @change="handleDaysChange">
    </days-input>

  </div>

</template>

<script>

  import OptionInput from '../common/option-input.vue';
  import TimeInput from '../common/time-input.vue';
  import DaysInput from '../common/days-input.vue';

  export default {

    name: 'start-filter',

    props: ['id', 'active', 'options'],

    components: {OptionInput, TimeInput, DaysInput},

    data() {

      return {

        operators: ['Before', 'Exactly', 'After'],

        changes: {...this.options}

      }

    },

    methods: {

      handleActiveToggle() {

        this.$store.dispatch('toggleFilter', this.id);

      },

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

      }

    }

  }

</script>

<style scoped>

  .start-filter {
    display: flex;
    align-items: center;
  }

  .start-filter > * + * {
    margin-left: 10px;
  }

</style>
