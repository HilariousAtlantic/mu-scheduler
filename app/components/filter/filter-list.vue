<template>

  <div class="filter-list">

    <h4>Time Filters</h4>

    <time-filter v-for="filter in timeFilters" :filter="filter"></time-filter>

    <h4>Class Filters</h4>

    <class-filter v-for="filter in classFilters" :filter="filter"></class-filter>

  </div>

</template>

<script>

  import TimeFilter from './time-filter.vue';
  import ClassFilter from './class-filter.vue';

  export default {

    name: 'filter-list',

    components: {TimeFilter, ClassFilter},

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


</style>
