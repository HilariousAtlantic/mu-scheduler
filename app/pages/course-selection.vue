<template>

  <div class="course-selection">

    <sidebar>

      <div class="sidebar-content">

        <h3>Course Selection</h3>

        <p>Welcome to the unofficial course scheduler for Miami University! Select the term and courses you would like to take.</p>

        <p>You can use the course filter to search for the course you want to select. You do not have to worry about using the right case or symbols.</p>

        <p>If you would like to know more about a course, click the <i class="fa fa-info-circle"></i> icon next the course you are interested in.</p>

        <p>Once you select a course, you can unselect it by clicking on it in the selected courses list.</p>

        <p>Once you select all of the courses you would like to take, click the Generate Schedules button. You will be taken to a screen where you can view all of the possible schedules.</p>

        <p>You can return this screen at any time by clicking on the Courses tab at the top of the window.</p>

      </div>

    </sidebar>

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

<style scoped>

  .course-selection {
    display: flex;
  }

  .sidebar {
    width: 300px;
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
