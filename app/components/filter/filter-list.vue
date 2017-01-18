<template>

  <div class="filter-list">

    <h4>Time Filters</h4>

    <time-filter></time-filter>

    <h4>Class Filters</h4>

    <class-filter></class-filter>

  </div>

</template>

<script>

  import TimeFilter from './time-filter.vue';
  import ClassFilter from './class-filter.vue';

  export default {

    name: 'filter-list',

    components: {TimeFilter, ClassFilter}

  }

</script>

<style scoped>


</style>
