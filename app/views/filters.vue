<template>

  <div class="filters-view">

    <header>

      <span class="title">Schedule Filters</span>

      <span>{{count}} Schedules Remain</span>

    </header>

    <filter-list></filter-list>

    <nav>

      <router-link to="/courses">

        <button type="button"><i class="fa fa-arrow-left"></i> Select Schedules</button>

      </router-link>

      <router-link to="/schedules">

        <button type="button">View Schedules <i class="fa fa-arrow-right"></i></button>

      </router-link>

    </nav>

  </div>

</template>

<script>

  import FilterList from '../components/filter/filter-list.vue';

  export default {

    name: 'filters-view',

    components: {FilterList},

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

    width: 90%;
    max-width: 1000px;
    margin: 20px auto;

  }

  header {

    display: flex;
    font-weight: 900;
    font-size: 1.25rem;
    margin-bottom: 20px;

  }

  .title {

    flex: 1;

  }

  nav {

    display: flex;
    justify-content: space-between;
    margin-top: 50px;

  }

  button {

    background: #0D47A1;
    color: #fff;
    padding: 10px 25px;
    outline: none;
    cursor: pointer;

  }

  .fa-arrow-left {

    margin-right: 25px;

  }

  .fa-arrow-right {

    margin-left: 25px;

  }

</style>
