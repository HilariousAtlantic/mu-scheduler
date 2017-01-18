<template>

  <div class="schedule-browser">

    <div class="browser-group">

      <button @click="decreaseIndex"><i class="fa fa-angle-left"></i></button>

      <span>{{$store.state.schedulesIndex+1}} of {{$store.state.schedules.length}}</span>

      <button @click="increaseIndex"><i class="fa fa-angle-right"></i></button>

    </div>

    <div class="browser-group">

      <button class="schedule-details">Details</button>

      <button class="schedule-select">Save</button>

      <button class="schedule-export">Export</button>

    </div>

  </div>

</template>

<script>

  export default {

    name: 'schedule-browser',

    methods: {

      increaseIndex() {

        this.$store.dispatch('increaseSchedulesIndex');

      },

      decreaseIndex() {

        this.$store.dispatch('decreaseSchedulesIndex');

      }

    }

  }

</script>

<style scoped>

  .schedule-browser {

    display: flex;
    justify-content: space-between;

  }

  button {

    border: 1px solid #ddd;
    padding: 10px 50px;
    background: #eee;
    outline: none;
    cursor: pointer;

    &:hover {

      background: #ddd;

    }

  }

  span {

    text-align: center;
    border: 1px solid #ddd;
    padding: 10px 75px;

  }

  .schedule-details {

    background: #1565C0;
    color: #fff;

    &:hover {

      background: #0D47A1;

    }

  }

  .schedule-select {

    background: #4CAF50;
    color: #fff;

    &:hover {

      background: #43A047;

    }

  }

  .schedule-export {

    background: #fff;
    color: #333;

    &:hover {

      background: #F5F5F5;

    }

  }

</style>
