<template>

  <div class="schedule-selection">

    <sidebar>

      <div class="sidebar-content">

        <h3>Schedule Selection</h3>

        <filter-list></filter-list>

      </div>

    </sidebar>

    <schedule-list></schedule-list>

  </div>

</template>

<script>

  import Sidebar from '../components/sidebar.vue';
  import FilterList from '../components/filter-list.vue';
  import ScheduleList from '../components/schedule/schedule-list.vue';

  export default {

    name: 'schedule-selection-page',

    components: {Sidebar, FilterList, ScheduleList}

  }

</script>

<style scoped>

  .schedule-selection {

    display: flex;

  }

  .sidebar {

    width: 350px;

  }

  .schedule-list {

    flex: 1;
    padding: 25px;

  }

  .sidebar-content {

    h3 {

      margin-top: 0;
      margin-bottom: 15px;

    }

    h4 {

      display: inline-block;
      margin: 0;

    }

  }

</style>
