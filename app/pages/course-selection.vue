<template>

  <div class="course-selection-page">

    <sidebar></sidebar>

    <course-selector></course-selector>

  </div>

</template>

<script>

  import Sidebar from '../components/sidebar.vue';
  import CourseSelector from '../components/course/course-selector.vue';

  export default {

    name: 'course-selection-page',

    components: {Sidebar, CourseSelector}

  }

</script>

<style lang="postcss" scoped>

  .course-selection-page {
    display: flex;
  }

  .sidebar {
    width: 300px;
  }

  .course-selector {
    flex: 1;
    padding: 25px;
  }

</style>
