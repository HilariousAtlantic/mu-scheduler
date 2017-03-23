<template>

  <div class="time-input">

    <input
      type="text"
      class="form-control"
      placeholder="HH:MM AM"
      v-model="time"
      @input="handleTimeChange"
    />

  </div>

</template>

<script>

  import {formatTime, getMinutes} from '../../lib/time';

  export default {

    name: 'time-input',

    props: ['defaultValue'],

    data() {

      return {

        time: formatTime(this.defaultValue, 'hh:mm P')

      }

    },

    methods: {

      handleTimeChange(event) {

        this.time = event.target.value.toUpperCase().substring(0, 8);

        if (/\d+:\d+ (AM|PM)/.test(this.time)) {

          console.log(this.time, getMinutes(this.time));

          this.$emit('change', getMinutes(this.time));

        }

      }

    }

  }

</script>

<style scoped>

  input {
    height: 34px;
    width: 75px;
    text-align: center;
  }

</style>
