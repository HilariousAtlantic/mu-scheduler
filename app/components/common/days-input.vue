<template>

  <div class="days-input btn-group tooltipped tooltipped-n"
    :aria-label="formattedDays">

    <button
      class="btn"
      v-for="day in days"
      :class="{selected: isSelected(day)}"
      @click="handleDayChange(day)"
    >{{day}}</button>

  </div>

</template>

<script>

  import {formatDayList} from '../../lib/days';

  export default {

    name: 'days-input',

    props: ['defaultValue'],

    data() {

      return {

        days: 'MTWRF',

        selectedDays: [...this.defaultValue]

      }

    },

    computed: {

      formattedDays() {

        return formatDayList(this.selectedDays);

      }

    },

    methods: {

      handleDayChange(day) {

        let index = this.selectedDays.indexOf(day);

        if (index === -1) {

          this.selectedDays = this.selectedDays.concat(day);

        } else {

          this.selectedDays.splice(index, 1);

        }

        this.$emit('change', this.selectedDays);

      },

      isSelected(day) {

        return this.selectedDays.indexOf(day) !== -1;

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
