<template>

  <div class="courses-view">

    <course-search></course-search>

    <course-list></course-list>

    <button
      class="btn btn-block btn-primary"
      @click="generateSchedules"
    >Generate Schedules</button>

  </div>

</template>

<script>

  import CourseSearch from '../components/course/course-search.vue';
  import CourseList from '../components/course/course-list.vue';

  export default {

    name: 'courses-view',

    components: {CourseSearch, CourseList},

    methods: {

      generateSchedules() {

        this.$store.dispatch('generateSchedules').then(schedules => {

          if (schedules) this.$router.push('schedules');

        });

      }

    }

  }

</script>

<style scoped>

  .course-list {
    margin: 10px 0;
  }

  .generate-button {
    font-weight: normal;
  }

</style>
