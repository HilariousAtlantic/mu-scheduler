<template>

  <div class="course-list">

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

    components: {Course}

  }

</script>

<style scoped>

</style>
