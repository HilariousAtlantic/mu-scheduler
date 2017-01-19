<template>

  <div class="filter-list">

    <h4>Time Filters</h4>

    <time-filter
      v-for="filter in timeFilters"
      :options="filter.options"
      :active="filter.active"
    ></time-filter>

    <h4>Class Filters</h4>

    <class-filter
      v-for="filter in classFilters"
      :options="filter.options"
      :active="filter.active"
    ></class-filter>

  </div>

</template>

<script>

  import TimeFilter from './time-filter.vue';
  import ClassFilter from './class-filter.vue';

  export default {

    name: 'filter-list',

    components: {TimeFilter, ClassFilter},

    data() {

      return {

        filters: [

            {type: 'time', options: {operator: 'Start After', time: '10:00 AM', days: ['M', 'W', 'F']}, active: true},

            {type: 'time', options: {operator: 'End Before', time: '4:00 PM', days: ['T', 'R']}, active: false},

            {type: 'class', options: {operator: 'At Most', amount: '2', days: ['M']}, active: false},

            {type: 'class', options: {operator: 'Exactly', amount: '3', days: ['M', 'T', 'W', 'R', 'F', 'S']}, active: true}

          ]

      }

    },

    computed: {

      timeFilters() {

        return this.filters.filter(({type}) => type === 'time');

      },

      classFilters() {

        return this.filters.filter(({type}) => type === 'class');

      }

    }

  }

</script>

<style scoped>


</style>
