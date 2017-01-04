<template>

  <div class="app">

    <h1>Schedule Buddy</h1>

    <h2>Select courses you would like to take this semester</h2>

    <ul>

      <li v-for="course in courses">{{course.subject}} {{course.number}} - {{course.name}}</li>

    </ul>

  </div>

</template>

<script>

  import axios from 'axios';

  export default {

    name: 'app',

    data() {

      return {
        courses: []
      }

    },

    mounted() {

      axios.get('/courses')
        .then(response => this.courses = response.data.slice(0, 30))

    }

  }

</script>

<style>

  h1, h2 {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
    font-weight: normal;
  }

</style>
