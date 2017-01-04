<template>

  <div class="app">

    <h1>Welcome to Schedule Buddy</h1>
    <h2>Select the courses you would like to take this semester</h2>

    <div class="course-selection">

      <div class="course-selection-column">

        <div class="course-search">
          <span>Spring 2017</span>
          <input type="text" placeholder="Filter Courses" v-model="filter" />
        </div>

        <ul class="course-list">
          <li v-for="course in filteredCourses" v-on:click="handleAddCourse">{{course}}</li>
        </ul>

      </div>

      <div class="course-selection-column">

        <ul class="selected-courses">
          <li>Selected Courses</li>
          <li v-for="course in selectedCourses" v-on:click="handleRemoveCourse">{{course}}</li>
        </ul>

      </div>

    </div>

  </div>

</template>

<script>

  import axios from 'axios';

  export default {

    name: 'app',

    data() {

      return {
        courses: [],
        selectedCourses: [],
        filter: ''
      }

    },

    mounted() {

      axios.get('/courses').then(response => this.courses = response.data);

    },

    methods: {

      handleAddCourse(event) {

        let course = event.target.innerText;

        if (this.selectedCourses.indexOf(course) < 0) {
          this.selectedCourses.push(course);
        }

      },

      handleRemoveCourse(event) {

        let course = event.target.innerText;

        this.selectedCourses.splice(this.selectedCourses.indexOf(course), 1);

      }


    },

    computed: {

      filteredCourses() {

        return this.courses
          .map(course => course.subject + ' ' + course.number + ' - ' + course.name)
          .filter(course => course.toLowerCase().indexOf(this.filter.toLowerCase()) > -1)

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
  }

  .course-selection {
    text-align: left;
    display: flex;
  }

  .course-selection-column {
    flex: 1;
    padding: 20px;
  }

  .course-selection ul {
    padding: 0;
    margin: 0;
    border: 1px solid #ddd;
    max-height: 500px;
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

  .selected-courses li:first-of-type {
    background: #eee;
    cursor: default;
  }

  .course-search {
    display: flex;
    align-items: center;
    background: #eee;
    border: 1px solid #ddd;
    margin-bottom: 5px;
  }

  .course-search span {
    padding: 0 10px;
  }

  .course-search input {
    flex: 1;
    border: none;
    font-size: 1rem;
    padding: 10px;
    outline: none;
  }

</style>
