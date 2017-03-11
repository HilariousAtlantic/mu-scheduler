<template>

  <div class="schedules-view">

    <filter-list></filter-list>

    <schedule-list></schedule-list>

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


</style>
