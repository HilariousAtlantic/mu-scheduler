<template>

  <div class="break-filter filter">

    <time-input
      :defaultValue="changes.start"
      @change="handleStartChange"
    ></time-input>

    <time-input
      :defaultValue="changes.finish"
      @change="handleFinishChange"
    ></time-input>

    <days-input
      :defaultValue="changes.days"
      @change="handleDaysChange"
    ></days-input>

    <button class="btn btn-danger" @click="handleFilterDelete">Remove</button>

  </div>

</template>

<script>

  import TimeInput from '../common/time-input.vue';
  import DaysInput from '../common/days-input.vue';

  export default {

    name: 'break-filter',

    props: ['id', 'options'],

    components: {TimeInput, DaysInput},

    data() {

      return {

        changes: {...this.options}

      }

    },

    methods: {

      handleStartChange(start) {

        this.changes.start = start;
        this.submitChanges();

      },

      handleFinishChange(finish) {

        this.changes.finish = finish;
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

  .break-filter {
    display: flex;
    align-items: center;
  }

  .break-filter > * + * {
    margin-left: 10px;
  }

</style>
