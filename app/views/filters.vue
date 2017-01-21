<template>

  <div class="filters-view">

    <header>Schedule Filters</header>

    <filter-list></filter-list>

  </div>

</template>

<script>

  import FilterList from '../components/filter/filter-list.vue';

  export default {

    name: 'filters-view',

    components: {FilterList}

  }

</script>

<style scoped>

  .filters-view {

    width: 90%;
    max-width: 1000px;
    margin: 20px auto;

  }

  header {

    font-weight: 900;
    font-size: 1.25rem;
    margin-bottom: 20px;

  }

</style>
