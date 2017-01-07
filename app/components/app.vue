<template>

  <div class="app">

    <navbar></navbar>

    <div class="page">

      <sidebar></sidebar>

      <course-selection></course-selection>

    </div>


  </div>

</template>

<script>

  import Navbar from './navbar.vue';
  import Sidebar from './sidebar.vue';
  import CourseSelection from './course-selection.vue';

  export default {

    name: 'app',

    components: {Navbar, Sidebar, CourseSelection}

  }

</script>

<style scoped>

  .app {

    height: 100vh;
    display: flex;
    flex-direction: column;

  }

  .page {

    flex: 1;
    display: flex;

  }

  .sidebar {

    width: 300px;

  }

  .course-selection {

    padding: 25px;
    flex: 1;

  }

</style>
