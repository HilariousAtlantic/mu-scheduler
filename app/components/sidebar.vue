<template>

  <div class="sidebar">

    <div class="sidebar-content">

      <h3>Course Selection</h3>

      <p>Welcome to Schedule Buddy! Select the semester and courses you would like to take.</p>

      <p>You can use the course filter to search for the course you want to select. You do not have to worry about using the right case or symbols.</p>

      <p>If you would like to know more about a course, click the <i class="fa fa-info-circle"></i> icon next the course you are interested in.</p>

      <p>Once you select a course, you can unselect it by clicking on it in the selected courses list.</p>

      <p>Once you select all of the courses you would like to take, click the Generate Schedules button. You will be taken to a screen where you can view all of the possible schedules.</p>

      <p>You can return this screen at any time by clicking on the Courses tab at the top of the window.</p>

    </div>

    <div class="sidebar-footer">

      <a href="https://github.com/hilariousatlantic">Developed by Hilarious Atlantic</a>

    </div>


  </div>

</template>

<script>

  export default {

    name: 'sidebar'

  }

</script>

<style lang="postcss" scoped>

  .sidebar {

    display: flex;
    flex-direction: column;
    background: #efefef;
    border-right: 1px solid #ddd;

  }

  .sidebar-content {

    padding: 20px;
    flex: 1;
    overflow-y: auto;

  }

  .sidebar-footer {

    text-align: center;
    padding: 10px 20px;
    border-top: 1px solid #ddd;

  }

  h3 {

    margin-top: 0;

  }

  a {

    color: #888;
    text-decoration: none;

    &:hover {

      color: #555;

    }

  }

</style>
