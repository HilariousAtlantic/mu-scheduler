<template>

  <div class="schedule-filter">

    <input type="checkbox" :checked="active">

    <span>{{text}}</span>

  </div>

</template>

<script>

  export default {

    name: 'schedule-filter',

    props: ['text', 'active']

  }

</script>

<style scoped>

  .schedule-filter {

    display: flex;
    align-items: center;
    margin-bottom: 10px;

  }

  input {

    margin-right: 5px;

  }

  span {

    flex: 1;

  }

</style>
