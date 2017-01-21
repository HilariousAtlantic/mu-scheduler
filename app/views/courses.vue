<template>

  <div class="courses-view">

    <course-search></course-search>

    <selected-courses></selected-courses>

    <router-link to="/filters">

      <button type="button" class="schedules-generate" @click="generateSchedules">Generate Schedules</button>

    </router-link>

  </div>

</template>

<script>

  import CourseSearch from '../components/course/course-search.vue';
  import SelectedCourses from '../components/course/selected-courses.vue';

  export default {

    name: 'courses-view',

    components: {CourseSearch, SelectedCourses},

    methods: {

      generateSchedules() {

        this.$store.dispatch('generateSchedules');

      }

    }

  }

</script>

<style scoped>

  .courses-view {

    width: 90%;
    max-width: 1000px;
    margin: 20px auto;

  }

  .selected-courses {

    margin: 20px 0;

  }

  .schedules-generate {

    width: 100%;
    padding: 15px;
    border: 1px solid #ddd;
    background: #4CAF50;
    color: #fff;
    outline: none;
    cursor: pointer;

  }

</style>
