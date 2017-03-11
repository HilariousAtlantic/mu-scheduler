<template>

  <div class="course-list">

    <div v-if="showInstructions" class="blankslate">
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

  import Course from './course.vue';

  export default {

    name: 'course-list',

    components: {Course},

    computed: {

      showInstructions() {

        return this.$store.state.selectedCourses.length === 0;

      }

    }

  }

</script>

<style scoped>

  .course-search {

    margin-bottom: 20px;

  }

</style>
