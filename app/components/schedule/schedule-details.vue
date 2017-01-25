<template>

  <div class="schedule-details">

    <div class="details-header">

      <span>Schedule Details</span>

    </div>

    <div class="details-body">

      <span>Average GPA: {{gpa}}</span><br>

      <span>Online Courses: {{onlineCourses}}</span>

    </div>

  </div>

</template>

<script>

  import {formatTime} from '../../lib/time';
  import {sortDays} from '../../lib/days';

  export default {

    name: 'schedule-details',

    props: ['courses', 'gpa'],

    computed: {

      onlineCourses() {

        return this.courses.filter(course => {

          for (let meet of course.meets) {

            if (meet.location === 'WEB') {

              return true;

            }

          }

          return false;

        }).map(course => course.name).join(', ');

      }

    }

  }

</script>

<style scoped>


  .schedule-details {

    border: 1px solid #ddd;

  }

  .details-header {

    padding: 10px;
    background: #eee;
    border-bottom: 1px solid #ddd;

  }

  .details-body {

    padding: 10px

  }

</style>
