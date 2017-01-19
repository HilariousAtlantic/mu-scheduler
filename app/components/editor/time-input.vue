<template>

  <div class="time-input">

    <input type="text" v-model="time">

    <ul v-if="suggestedTimes.length > 1" class="suggestion-list">

      <li v-for="time in suggestedTimes" @click="selectTime(time)">{{time}}</li>

    </ul>

  </div>

</template>

<script>

  import {getSuggestedTimes} from '../lib/time';

  export default {

    name: 'time-input',

    props: ['step', 'defaultTime'],

    data() {

      return {

        time: this.defaultTime + '' || ''

      }

    },

    methods: {

      selectTime(time) {

        this.time = time;

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
