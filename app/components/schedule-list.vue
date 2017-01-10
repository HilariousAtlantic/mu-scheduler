<template>

  <div class="schedule-list">

    <schedule></schedule>

  </div>

</template>

<script>

  import Schedule from './schedule.vue';

  export default {

    name: 'schedule-list',

    components: {Schedule}

  }

</script>

<style lang="postcss" scoped>

  .schedule-list {

    display: flex;

  }

  .schedule {

    flex: 1;

  }

</style>
