<template>

  <div class="course-selection">

    <course-selector></course-selector>

  </div>

</template>

<script>

  import CourseSelector from '../components/course/course-selector.vue';

  export default {

    name: 'course-selection-page',

    components: {CourseSelector}

  }

</script>

<style scoped>

  .course-selection {
    display: flex;
  }

  .course-selector {
    flex: 1;
    padding: 25px;
  }

  .sidebar-content {

    h3 {

      margin-top: 0;

    }

  }

</style>
