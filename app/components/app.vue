<template>

  <div class="app">

    <h1>Welcome to Schedule Buddy</h1>
    <h2>Select the semester and the courses you would like to take</h2>

    <div class="course-selection">

      <div class="course-selection-column">

        <div class="course-search">

          <span @click="toggleSemesterList">{{selectedSemester}}<i class="fa fa-caret-down"></i></span>
          <input type="text" placeholder="Filter Courses" autocomplete="off" v-model="filter" />

          <ul class="semester-list" v-if="showSemesterList">
            <li v-for="semester in semesters" @click="selectSemester(semester)">{{semester}}</li>
          </ul>

        </div>

        <ul class="course-list">
          <li v-for="course in filteredCourses" @click="selectCourse(course)">{{course}}</li>
        </ul>

      </div>

      <div class="course-selection-column">

        <ul class="selected-courses">
          <li>Selected Courses for {{selectedSemester}}</li>
          <li v-for="course in selectedCourses" @click="unselectCourse(course)">{{course}}</li>
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
        semesters: ['Fall 2016', 'Winter 2017', 'Spring 2017', 'Summer 2017'],
        courses: [],
        selectedCourses: [],
        selectedSemester: 'Spring 2017',
        filter: '',
        showSemesterList: false
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
        this.showSemesterList = false;

      },

      toggleSemesterList() {

        this.showSemesterList = !this.showSemesterList;

      }


    },

    computed: {

      filteredCourses() {

        let term = this.filter.toLowerCase().replace(/\W+/g, '');

        return this.courses
          .map(course => course.subject + ' ' + course.number + ' - ' + course.name)
          .filter(course => course.toLowerCase().replace(/\W+/g, '').indexOf(term) > -1)

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
    position: relative;
    background: #eee;
    border: 1px solid #ddd;
    margin-bottom: 5px;
  }

  .semester-list {
    left: -1px;
    top: 100%;
    margin: 0;
    padding: 0;
    position: absolute;
  }

  .semester-list li {
    width: 7.5rem;
    background: #fff;
    list-style: none;
  }

  .course-search span {
    width: 7.5rem;
    padding: 10px;
    cursor: pointer;
  }

  .course-search i {
    float: right;
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
    padding: 15px;
    border: 1px solid #ddd;
    background: #66BB6A;
    color: #fff;
    margin: 20px;
  }

</style>
