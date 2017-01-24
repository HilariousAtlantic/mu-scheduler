<template>

  <div class="time-input">

    <input type="text" v-model="time" @focus="focused = true" @blur="focused = false">

    <ul v-if="showSuggestions" class="suggestion-list">

      <li v-for="time in times" @mousedown="selectTime(time)">{{time}}</li>

    </ul>

  </div>

</template>

<script>

  import {getSuggestedTimes, toTime, toMinutes} from '../../lib/time';

  export default {

    name: 'time-input',

    props: ['step', 'defaultTime'],

    data() {

      return {

        time: toTime(this.defaultTime || 600),

        focused: false

      }

    },

    methods: {

      selectTime(time) {

        this.time = time;

        this.$emit('change', toMinutes(time));

      }

    },

    computed: {

      times() {

        return getSuggestedTimes(this.step, this.time).slice(0, 4);

      },

      showSuggestions() {

        return this.time.trim().length > 0 && this.focused;

      }

    }

  }

</script>

<style scoped>

  .time-input {

    display: inline-block;
    position: relative;

  }

  input {

    background: #fff;
    border: 1px solid #ddd;
    padding: 10px;
    width: 80px;
    outline: none;

  }

  ul {

    margin: 0;
    padding: 0;
    width: 100px;
    position: absolute;
    border: 1px solid #ddd;
    border-top: none;
    background: #fff;
    cursor: pointer;

  }

  li {

    list-style: none;
    padding: 10px;

    &:hover {

      background: #f5f5f5;

    }

  }

</style>
