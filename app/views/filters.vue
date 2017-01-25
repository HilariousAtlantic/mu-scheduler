<template>

  <div class="filters-view">

    <double-header>

      <span slot="left">Schedule Filters</span>

      <span slot="right">{{count}} Schedules Remain</span>

    </double-header>

    <filter-list></filter-list>

    <router-link to="schedules">

      <button type="button" class="schedules-view">View Schedules</button>

    </router-link>

  </div>

</template>

<script>

  import DoubleHeader from '../components/common/double-header.vue';
  import FilterList from '../components/filter/filter-list.vue';

  export default {

    name: 'filters-view',

    components: {DoubleHeader, FilterList},

    computed: {

      count() {

        let filtered = this.$store.getters.filteredSchedules.length;

        let total = this.$store.state.schedules.length;

        return filtered + ' of ' + total;

      }

    }

  }

</script>

<style scoped>

  .filters-view {

  }

  .double-header {

    font-weight: 900;
    font-size: 1.25rem;

  }

  .filter-list {
    margin: 20px 0;
  }

  .schedules-view {

    width: 100%;
    padding: 15px;
    border: 1px solid #ddd;
    background: #1565C0;
    color: #fff;
    outline: none;
    cursor: pointer;

    &:hover {

      background: #0D47A1;

    }

  }

</style>
