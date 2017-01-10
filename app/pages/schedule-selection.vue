<template>

  <div class="schedule-selection">

    <sidebar></sidebar>

    <schedule-list></schedule-list>

  </div>

</template>

<script>

  import Sidebar from '../components/sidebar.vue';
  import ScheduleList from '../components/schedule-list.vue';

  export default {

    name: 'schedule-selection-page',

    components: {Sidebar, ScheduleList}

  }

</script>

<style lang="postcss" scoped>

  .schedule-selection {

    display: flex;

  }

  .sidebar {

    width: 300px;

  }

  .schedule-list {

    flex: 1;
    padding: 25px;

  }


</style>
