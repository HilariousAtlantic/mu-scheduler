<template>

  <div class="app">

    <h1>Welcome to Schedule Buddy</h1>
    <h2>Select the courses you would like to take this semester</h2>

    <div class="course-selection">

      <div class="course-selection-column">

        <div class="course-search">
          <span>{{semester}}</span>
          <input type="text" placeholder="Filter Courses" autocomplete="off" v-model="filter" />
        </div>

        <ul class="course-list">
          <li v-for="course in filteredCourses" v-on:click="handleAddCourse">{{course}}</li>
        </ul>

      </div>

      <div class="course-selection-column">

        <ul class="selected-courses">
          <li>Selected Courses for {{semester}}</li>
          <li v-for="course in selectedCourses" v-on:click="handleRemoveCourse">{{course}}</li>
        </ul>

      </div>

    </div>

    <button type="button">Generate Schedules</button>

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
        semester: 'Spring 2017',
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

  .selected-courses li:first-of-type {
    background: #eee;
    cursor: default;
  }

  .selected-courses {
    max-height: 545px;
  }

  .course-list {
    max-height: 500px;
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
    border-left: 1px solid #ddd;
    font-size: 1rem;
    padding: 10px;
    outline: none;
  }

  button {
    width: 100%;
    max-width: 1000px;
    padding: 10px;
    border: 1px solid #ddd;
    background: #66BB6A;
    color: #fff;
    margin: 20px;
  }

</style>
