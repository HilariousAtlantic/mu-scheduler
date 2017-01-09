<template>

  <div class="course-selector">

    <course-search></course-search>

    <div class="course-list-container">

      <course-list
        :header="'Course List' + selectedSemester"
        :courses="filteredCourses"
        dispatch="selectCourse"
      ></course-list>

      <course-list
        :header="'Selected Courses' + selectedSemester"
        :courses="$store.state.selectedCourses"
        dispatch="deselectCourse"
      ></course-list>

    </div>

    <button type="button">Generate Schedules</button>

  </div>

</template>

<script>

  import CourseSearch from './course-search.vue';
  import CourseList from './course-list.vue';

  export default {

    name: 'course-selector',

    components: {CourseSearch, CourseList},

    computed: {

      selectedSemester() {

        let semester = this.$store.state.selectedSemester;

        return semester ? ' for ' + semester.name : '';

      },

      filteredCourses() {

        let semester = this.$store.state.selectedSemester;

        if (!semester) {
          return [];
        }

        let courses = this.$store.state.courses[semester.id];

        if (!courses) {
          return [];
        }

        let trim = (term) => term.toLowerCase().replace(/\W+/g, '');

        return courses.filter(({subject, number, name}) =>
          trim(subject+number+name).indexOf(trim(this.$store.state.coursesFilter)) != -1
        );

      }

    }

  }

</script>

<style lang="postcss" scoped>

  .course-selector {

    display: flex;
    flex-direction: column;
    font-family: sans-serif;
    color: #333;

  }

  .course-list-container {

    flex: 1;
    display: flex;
    margin: 10px 0;

  }

  .course-list {

    flex: 1;

    &:first-of-type {

      margin-right: 10px;

    }

  }

  button {

    width: 100%;
    padding: 15px;
    border: 1px solid #ddd;
    background: #66BB6A;
    color: #fff;

  }

</style>
