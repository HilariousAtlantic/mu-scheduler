<template>

  <div class="course-selector">

    <course-search></course-search>

    <div class="course-list-container">

      <filtered-courses></filtered-courses>

      <selected-courses></selected-courses>

    </div>

    <button type="button">Generate Schedules</button>

  </div>

</template>

<script>

  import CourseSearch from './course-search.vue';
  import FilteredCourses from './filtered-courses.vue';
  import SelectedCourses from './selected-courses.vue';

  export default {

    name: 'course-selector',

    components: {CourseSearch, FilteredCourses, SelectedCourses},

    computed: {

      selectedSemester() {

        let semester = this.$store.state.selectedSemester;

        return semester ? ' for ' + semester.name : '';

      },

      filteredCourses() {

        let trim = (term) => term.toLowerCase().replace(/\W+/g, '');

        return this.$store.state.courses.filter(({subject, number, name}) =>
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

  .filtered-courses {

    flex: 1;
    margin-right: 10px;

  }

  .selected-courses {

    flex: 1;

  }

  button {

    width: 100%;
    padding: 15px;
    border: 1px solid #ddd;
    background: #66BB6A;
    color: #fff;

  }

</style>
