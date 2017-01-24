<template>

  <div class="schedule-browser">

    <button @click="$emit('prev')">Previous Schedule</button>

    <button @click="$emit('next')">Next Schedule</button>

  </div>

</template>

<script>

  export default {

    name: 'schedule-browser'

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

</style>
