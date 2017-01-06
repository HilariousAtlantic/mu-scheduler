<template>

  <div class="app">

    <h1>Welcome to Schedule Buddy</h1>
    <h2>Select the semester and courses you would like to take</h2>

    <course-search
      :semesters="semesters"
      :selectedSemester="selectedSemester"
      @selectSemester="onSemesterSelect"
      @changeFilter="onFilterChange"
    ></course-search>

    <div class="course-selection">

      <course-list
        :header="'Course List for ' + selectedSemester"
        :courses="filteredCourses"
        @clickCourse="onCourseSelect"
      ></course-list>

      <course-list
        :header="'Selected Courses for ' + selectedSemester"
        :courses="selectedCourses"
        @clickCourse="onCourseUnselect"
      ></course-list>

    </div>

    <button type="button">Generate Schedules</button>

  </div>

</template>

<script>

  import axios from 'axios';

  import CourseSearch from './course-search.vue';
  import CourseList from './course-list.vue';

  export default {

    name: 'app',

    components: {CourseSearch, CourseList},

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
    display: flex;
    flex-direction: column;
    width: 90%;
    height: 100vh;
    max-width: 1000px;
    margin: 0 auto;
    font-family: sans-serif;
    color: #333;
  }

  h1, h2 {
    margin: 10px 0;
    text-align: center;
  }

  .course-selection {
    display: flex;
    align-items: flex-start;
    flex: 1;
    margin: 10px 0;
  }

  .course-list:first-of-type {

    margin-right: 10px;

  }

  button {
    width: 100%;
    padding: 15px;
    border: 1px solid #ddd;
    background: #66BB6A;
    color: #fff;
    margin-bottom: 10px;
  }

</style>
