<template>

  <div class="selected-courses">

    <header>Selected Courses {{selectedSemester}}</header>

    <ul>

      <li v-for="course in $store.state.selectedCourses" @click="onCourseClick(course)">

        {{course.subject}} {{course.number}} - {{course.name}}
        <i class="course-info fa fa-info-circle"></i>

      </li>

    </ul>

  </div>



</template>

<script>

  export default {

    name: 'selected-courses',

    methods: {

      onCourseClick(course) {

        this.$store.dispatch('deselectCourse', course);

      }

    },

    computed: {

      selectedSemester() {

        let semester = this.$store.state.selectedSemester;

        return semester ? 'for ' + semester.name : '';

      }

    }

  }

</script>

<style lang="postcss" scoped>

  .selected-courses {

    display: flex;
    flex-direction: column;

  }

  header {

    background: #eee;
    border: 1px solid #ddd;
    padding: 10px;

  }

  ul {

    padding: 0;
    margin: 0;
    flex: 1;
    overflow-y: auto;
    border: 1px solid #ddd;
    border-top: none;

  }

  li {

    list-style: none;
    padding: 10px;
    cursor: pointer;

    &:hover {

      background: #f5f5f5;

      .course-info {

        display: inline-block;

      }

    }

  }

  .course-info {

    display: none;
    float: right;

  }

  .loader-wrapper {

    flex: 1;
    display: flex;
    justify-content: center;
    align-items: center;
    border: 1px solid #ddd;
    border-top: none;

  }

  .loader {

    width: 50px;
    height: 50px;
    border: 5px solid #ddd;
    border-bottom: 5px solid #444;
    border-radius: 90px;
    animation: spin 1s linear infinite;

  }

  @keyframes spin {

    to {

      transform: rotate(360deg);

    }

  }

</style>
