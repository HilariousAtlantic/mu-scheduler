<template>

  <div class="days-input btn-group">

    <button
      class="btn"
      v-for="day in days"
      @click="handleDayChange(day)"
    >{{day}}</button>

  </div>

</template>

<script>

  export default {

    name: 'days-input',

    props: ['defaultValue'],

    data() {

      return {

        days: 'MTWRF',

        selectedDays: [...this.defaultValue]

      }

    },

    methods: {

      handleDayChange(day) {

        let index = this.selectedDays.indexOf(day);

        if (index === -1) {

          this.selectedDays = this.selectedDays.concat(day)
            .sort((a, b) => this.days.indexOf(a)-this.days.indexOf(b));

        } else {

          this.selectedDays.splice(index, 1);

        }

        this.$emit('change', this.selectedDays);

      }

    }

  }

</script>

<style scoped>

  button {
    width: 35px;
    padding: 6px 0;
    text-align: center;
  }

</style>
