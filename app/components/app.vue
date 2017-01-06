<template>

  <div class="app">

    <h1>Welcome to Schedule Buddy</h1>
    <h2>Select the semester and the courses you would like to take</h2>

    <div class="course-selection">

      <div class="course-selection-column">

        <course-search
          :semesters="semesters"
          :selectedSemester="selectedSemester"
          @selectSemester="onSemesterSelect"
          @changeFilter="onFilterChange"
        ></course-search>

        <course-list
          :courses="filteredCourses"
          @selectCourse="onCourseSelect"
        ></course-list>

      </div>

      <div class="course-selection-column">

        <selected-courses
          :semester="selectedSemester"
          :courses="selectedCourses"
          @unselectCourse="onCourseUnselect"
        ></selected-courses>

      </div>

    </div>

    <button type="button">Generate Schedules</button>

  </div>

</template>

<script>

  import axios from 'axios';

  import CourseSearch from './course-search.vue';
  import CourseList from './course-list.vue';
  import SelectedCourses from './selected-courses.vue';

  export default {

    name: 'app',

    components: {CourseSearch, CourseList, SelectedCourses},

    data() {

      return {
        semesters: ['Fall 2016', 'Winter 2017', 'Spring 2017', 'Summer 2017'],
        courses: [],
        selectedCourses: [],
        selectedSemester: 'Spring 2017',
        filter: ''
      }

    },

    mounted() {

      axios.get('/courses').then(response => this.courses = response.data);

    },

    methods: {

      onCourseSelect(course) {

        if (this.selectedCourses.indexOf(course) == -1) {
          this.selectedCourses.push(course);
        }

      },

      onCourseUnselect(course) {

        let index = this.selectedCourses.indexOf(course);

        if (index != -1) {
          this.selectedCourses.splice(index, 1);
        }

      },

      onSemesterSelect(semester) {

        this.selectedSemester = semester;

      },

      onFilterChange(filter) {

        this.filter = filter;

      }

    },

    computed: {

      filteredCourses() {

        let trim = (term) => term.toLowerCase().replace(/\W+/g, '');

        return this.courses.filter(({subject, number, name}) =>
          trim(subject+number+name).indexOf(trim(this.filter)) != -1
        );

      }

    }

  }

</script>

<style scoped>

  .app {
    font-family: sans-serif;
    width: 90%;
    max-width: 1000px;
    margin: 0 auto;
    text-align: center;
    color: #333;
  }

  .course-selection {
    text-align: left;
    display: flex;
  }

  .course-selection-column {
    flex: 1;
    padding: 0 20px;
  }

  button {
    width: 100%;
    max-width: 1000px;
    padding: 15px;
    border: 1px solid #ddd;
    background: #66BB6A;
    color: #fff;
    margin: 20px;
  }

</style>
