<template>

  <div class="schedule-filter">

    <input type="checkbox" :checked="active" @change="$emit('toggle')">

    <span @click="$emit('edit')">{{text}}</span>

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

    height: 1rem;
    margin-right: 5px;

  }

  span {

    flex: 1;
    cursor: pointer;

    &:hover {

      background: rgba(0, 0, 0, .05);

    }

  }

</style>
