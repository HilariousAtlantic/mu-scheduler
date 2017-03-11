<template>

  <div class="courses-view">

    <course-search></course-search>

    <course-list></course-list>

    <router-link
      to="schedules"
      class="generate-button btn btn-block btn-primary"
    >Generate Schedules</router-link>

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

        this.$store.dispatch('generateSchedules');

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
