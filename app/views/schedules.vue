<template>

  <div class="schedules-view">

    <schedule-list v-if="$store.state.schedules.length"></schedule-list>

    <div v-else class="schedules-none">

      <span>No schedules generated</span>

      <router-link to="/courses">

        <button>Select Courses</button>

      </router-link>

    </div>

  </div>

</template>

<script>

  import FilterList from '../components/filter/filter-list.vue';
  import ScheduleList from '../components/schedule/schedule-list.vue';

  export default {

    name: 'schedules-view',

    components: {FilterList, ScheduleList}

  }

</script>

<style scoped>

  .schedules-view {

    display: flex;
    justify-content: center;
    align-items: center;
    width: 90%;
    max-width: 1000px;
    margin: 0 auto;

  }

  .schedule-list {

    flex: 1;
    padding: 25px;

  }

  .schedules-none {

    display: flex;
    flex-direction: column;
    align-items: center;

    span {

      font-size: 1.25rem;
      font-weight: 900;
      margin-bottom: 10px;

    }

    button {

      padding: 10px 50px;
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

</style>
