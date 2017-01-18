<template>

  <div v-if="$store.state.schedules.length === 0" class="schedules-none">

    <span>No schedules generated</span>

    <router-link to="/courses"><button>Select Courses</button></router-link>

  </div>

  <div v-else-if="$store.state.requestingSchedules" class="schedules-loading">

    <div class="loader"></div>

  </div>

  <div v-else class="schedule-selection">

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
  import FilterList from '../components/filter/filter-list.vue';
  import ScheduleList from '../components/schedule/schedule-list.vue';

  export default {

    name: 'schedule-selection-page',

    components: {Sidebar, FilterList, ScheduleList}

  }

</script>

<style scoped>

  .schedules-none {

    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;

    span {

      font-size: 2rem;
      font-weight: 900;
      margin-bottom: 10px;
      color: #ccc;
      text-transform: uppercase;

    }

    button {

      padding: 15px 150px;
      font-size: 1rem;
      border: 1px solid #ddd;
      background: #1565C0;
      color: #fff;
      outline: none;
      cursor: pointer;

      &:hover {

        background: #0D47A1;

      }

    }

  }

  .schedules-loading {

    display: flex;
    align-items: center;
    justify-content: center;

  }

  .loader {

    width: 100px;
    height: 100px;
    border: 10px solid #ddd;
    border-bottom: 10px solid #444;
    border-radius: 90px;
    animation: spin 1s linear infinite;

  }

  @keyframes spin {

    to {

      transform: rotate(360deg);

    }

  }

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
