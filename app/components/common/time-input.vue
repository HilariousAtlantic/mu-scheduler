<template>

  <input
    type="time"
    class="time-input form-control"
    v-model="time"
    @change="handleTimeChange"
  />

</template>

<script>

  import {formatTime, getMinutes} from '../../lib/time';

  export default {

    name: 'time-input',

    props: ['defaultValue'],

    data() {

      return {

        time: formatTime(this.defaultValue || 600, 'HH:mm')

      }

    },

    methods: {

      handleTimeChange() {

        this.$emit('change', getMinutes(this.time));

      }

    }

  }

</script>

<style scoped>

  input {
    height: 34px;
  }

</style>
