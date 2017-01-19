<template>

  <div class="days-input">

    <button v-for="day in days" :class="{selected: selectedDays.indexOf(day) != -1}" @click="toggleDay(day)">{{day}}</button>

  </div>

</template>

<script>

  export default {

    name: 'days-input',

    props: ['days', 'defaultDays'],

    data() {

      return {

        selectedDays: this.defaultDays || []

      }

    },

    methods: {

      toggleDay(day) {

        let index = this.selectedDays.indexOf(day);

        if (index === -1) {

          this.selectedDays.push(day);

        } else {

          this.selectedDays.splice(index, 1);

        }

      }

    },

    computed: {

      suggestedTimes() {

        return getSuggestedTimes(this.step, this.time).slice(0, 4);

      }

    }

  }

</script>

<style scoped>

  .days-input {

    display: inline-flex;
    border: 1px solid #ddd;

  }

  button {

    flex: 1;
    background: #fff;
    border: none;
    border-right: 1px solid #ddd;
    padding: 10px;
    width: 40px;
    outline: none;
    cursor: pointer;

    &:last-of-type {

      border-right: none;

    }

  }

  .selected {

    background: #f5f5f5;

  }

</style>
