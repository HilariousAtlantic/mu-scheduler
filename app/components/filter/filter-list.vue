<template>

  <div class="filter-list">

    <double-header>

      <span slot="left">Time Filters</span>

      <button type="button" slot="right">

        <i class="fa fa-plus"></i> Add Filter

      </button>

    </double-header>

    <time-filter v-for="filter in timeFilters" :filter="filter"></time-filter>

    <double-header>

      <span slot="left">Class Filters</span>

      <button type="button" slot="right">

        <i class="fa fa-plus"></i> Add Filter

      </button>

    </double-header>

    <class-filter v-for="filter in classFilters" :filter="filter"></class-filter>

  </div>

</template>

<script>

  import DoubleHeader from '../common/double-header.vue'
  import TimeFilter from './time-filter.vue';
  import ClassFilter from './class-filter.vue';

  export default {

    name: 'filter-list',

    components: {DoubleHeader, TimeFilter, ClassFilter},

    computed: {

      timeFilters() {

        return this.$store.state.scheduleFilters.filter(({type}) => type === 'time');

      },

      classFilters() {

        return this.$store.state.scheduleFilters.filter(({type}) => type === 'class');

      }

    }

  }

</script>

<style scoped>

  .double-header {

    margin-top: 20px;
    margin-bottom: 10px;
    font-weight: 900;

  }

  button {

    padding: 0;
    background: #fff;
    border: none;
    outline: none;
    cursor: pointer;
    font-weight: inherit;

    &:hover {

      border-bottom: 2px solid #333;

    }

  }


</style>
