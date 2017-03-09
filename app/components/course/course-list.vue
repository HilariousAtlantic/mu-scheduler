<template>

  <div class="course-list">

    <course-search></course-search>

    <div v-if="$store.state.selectedCourses.length == 0" class="blankslate">
      <h2>No Courses Selected</h2>
      <p>Use the search bar to select the classes you would like to take</p>
    </div>

    <course
      v-for="course in $store.getters.selectedCourses"
      :course="course"
    ></course>

  </div>

</template>

<script>

  import CourseSearch from './course-search.vue';
  import Course from './course.vue';

  export default {

    name: 'course-list',

    components: {CourseSearch, Course},

    methods: {

      selectTerm(term) {

        this.$store.dispatch('selectTerm', term);

      },

      selectCourse(course) {

        this.$store.dispatch('selectCourse', course);

      }

    }

  }

</script>

<style scoped>

  .course-search {

    margin-bottom: 20px;

  }

</style>
