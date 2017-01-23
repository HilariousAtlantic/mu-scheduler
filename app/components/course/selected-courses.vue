<template>

  <div class="selected-courses">

    <detailed-course v-for="course in $store.getters.selectedCourses" :course="course"></detailed-course>

  </div>

</template>

<script>

  import DetailedCourse from './detailed-course.vue';

  export default {

    name: 'selected-courses',

    components: {DetailedCourse},

    methods: {

      deselectCourse(course) {

        this.$store.dispatch('deselectCourse', course);

      }

    }

  }

</script>

<style scoped>

  .selected-courses {

    display: flex;
    flex-direction: column;

  }

</style>
