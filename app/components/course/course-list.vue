<template>

  <div class="course-list">

    <course-search></course-search>

    <course v-for="course in $store.getters.selectedCourses" :course="course"></course>

  </div>

</template>

<script>

  import CourseSearch from './course-search.vue';
  import Course from './course.vue';

  export default {

    name: 'course-list',

    components: {CourseSearch, Course},

    methods: {

      deselectCourse(course) {

        this.$store.dispatch('deselectCourse', course);

      }

    }

  }

</script>

<style scoped>

  .course-search {

    margin-bottom: 10px;

  }

</style>
