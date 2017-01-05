<template>

  <div class="app">

    <h1>Welcome to Schedule Buddy</h1>
    <h2>Select the semester and the courses you would like to take</h2>

    <div class="course-selection">

      <div class="course-selection-column">

        <course-search :semesters="semesters" :selectedSemester="selectedSemester" @filter="setFilter" @semester="selectSemester"></course-search>

        <course-list :courses="filteredCourses" @select="selectCourse"></course-list>

      </div>

      <div class="course-selection-column">

        <selected-courses @unselect="unselectCourse" :semester="selectedSemester" :courses="selectedCourses"></selected-courses>

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

      selectCourse(course) {

        if (this.selectedCourses.indexOf(course) == -1) {
          this.selectedCourses.push(course);
        }

      },

      unselectCourse(course) {

        let index = this.selectedCourses.indexOf(course);

        if (index != -1) {
          this.selectedCourses.splice(index, 1);
        }

      },

      selectSemester(semester) {

        this.selectedSemester = semester;

      },

      setFilter(filter) {

        this.filter = filter;

      }

    },

    computed: {

      filteredCourses() {

        let trim = (term) => term.toLowerCase().replace(/\W+/g, '');

        return this.courses
          .map(course => course.subject + ' ' + course.number + ' - ' + course.name)
          .filter(course => trim(course).indexOf(trim(this.filter)) > -1)

      }

    }

  }

</script>

<style>

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

  .course-selection ul {
    padding: 0;
    margin: 0;
    border: 1px solid #ddd;
    overflow-y: scroll;
  }

  .course-selection li {
    list-style: none;
    padding: 10px;
    border-top: 1px solid #ddd;
    cursor: pointer;
  }

  .course-selection li:hover {
    background: #f5f5f5;
  }

  .course-selection li:first-of-type {
    border: none;
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
