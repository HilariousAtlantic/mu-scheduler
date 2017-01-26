<template>

  <div class="filters-view">

    <instruction>

      <div>

        <h3>Creating Filters</h3>

        <p>There is a good chance that the number of possible schedules to choose from is too high to work with. Filters will help narrow down the choices to schedules that fit your needs.</p>

        <p>To create a filter, click the add filter button next to the type of filter you would like to create. You can edit the filter by clicking on it. Make sure to enable the filter, as they are disabled by default.</p>

      </div>

    </instruction>

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

  import Instruction from '../components/common/instruction.vue';
  import DoubleHeader from '../components/common/double-header.vue';
  import FilterList from '../components/filter/filter-list.vue';

  export default {

    name: 'filters-view',

    components: {Instruction, DoubleHeader, FilterList},

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

  .instruction {

    margin-bottom: 10px;

    h3 {

      margin-top: 0;

    }

    p {

      margin-top: 0;
      margin-bottom: 10px;

      &:last-of-type {

        margin-bottom: 0;

      }

    }

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
