<template>

  <div class="filter-list">

    <h4>Time Filters</h4>

    <time-filter
      v-for="filter in filters.time"
      :operator="filter.operator"
      :time="filter.time"
      :days="filter.days"
    ></time-filter>

    <h4>Class Filters</h4>

    <class-filter
      v-for="filter in filters.class"
      :operator="filter.operator"
      :amount="filter.amount"
      :days="filter.days"
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

        filters: {

          time: [

            {operator: 'Start After', time: '10:00 AM', days: 'Monday, Wednesday, and Friday', active: true},

            {operator: 'End Before', time: '4:00 PM', days: 'Tuesday and Thursday', active: false}

          ],

          class: [

            {operator: 'At Most', amount: '2', days: 'Monday', active: false},

            {operator: 'Exactly', amount: '3', days: 'Every Day', active: true}

          ]

        }

      }

    }

  }

</script>

<style scoped>


</style>
