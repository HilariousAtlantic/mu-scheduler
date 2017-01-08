<template>

  <div class="course-selector">

    <course-search
      :semesters="semesters"
      :selectedSemester="selectedSemester"
      @selectSemester="onSemesterSelect"
      @changeFilter="onFilterChange"
    ></course-search>

    <div class="course-list-container">

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

    name: 'course-selector',

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
