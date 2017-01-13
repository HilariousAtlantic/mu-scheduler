<template>

  <div class="schedule-selection">

    <sidebar>

      <div class="sidebar-content">

        <h3>Schedule Selection</h3>

        <div class="filter-group">

          <h4>Time Filters</h4>

          <button><i class="fa fa-plus"></i> Filter</button>

        </div>

        <div class="filter">

          <input type="checkbox">

          <span>I want to start after 10:00 AM on Monday, Wednesday, and Friday</span>

        </div>

        <div class="filter">

          <input type="checkbox">

          <span>I want to finish before 4:00 PM on Tuesday and Thursday</span>

        </div>

        <div class="filter-group">

          <h4>Class Filters</h4>

          <button><i class="fa fa-plus"></i> Filter</button>

        </div>

        <div class="filter">

          <input type="checkbox">

          <span>I want less than 3 classes on Monday, Wednesday, and Friday</span>

        </div>

        <div class="filter">

          <input type="checkbox">

          <span>I want more than 1 class on Tuesday and Thursday</span>

        </div>

      </div>

    </sidebar>

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

    .filter-group {

      margin: 15px 0;

      button {

        float: right;
        border: none;
        background: transparent;

      }

    }

    .filter {

      display: flex;
      margin-bottom: 10px;

      input {

        margin-right: 5px;

      }

      span {

        flex: 1;

      }

    }

  }


</style>
