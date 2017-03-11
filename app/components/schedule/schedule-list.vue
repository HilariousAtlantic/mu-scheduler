<template>

  <div class="schedule-list">

    <!---<schedule-details
      v-if="showDetails"
      :courses="schedule.courses"
      :gpa="schedule.gpa"
    ></schedule-details>

    <schedule-calendar
      :courses="schedule.courses"
      :start="schedule.start"
      :length="schedule.length"
    ></schedule-calendar>-->

  </div>

</template>

<script>

  import ScheduleDetails from './schedule-details.vue';
  import ScheduleCalendar from './schedule-calendar.vue';

  export default {

    name: 'schedule-list',

    props: ['index'],

    components: {ScheduleDetails, ScheduleCalendar},

    data() {

      return {

        showDetails: false

      }

    },

    computed: {

      schedule() {

        return this.$store.getters.filteredSchedules[this.index];

      },

      showHide() {

        return this.showDetails ? 'Hide' : 'More';

      }

    },

    methods: {

      toggleDetails() {

        this.showDetails = !this.showDetails;

      }

    }

  }

</script>

<style scoped>

</style>
