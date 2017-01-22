<template>

  <div class="schedule">

    <div class="schedule-head">

      <span v-for="day in days">{{day}}</span>

    </div>

    <div class="schedule-body" :style="scheduleHeight">

      <div class="day" v-for="day in days">

          <div class="course" v-for="course in coursesByDay[day]" :style="course.style">

            <div class="course-meta">

              <span>{{course.name}}</span>
              <span>{{course.start}} - {{course.end}} {{course.location}}</span>
              <span>{{course.instructor}}</span>

            </div>

          </div>

      </div>

    </div>

  </div>

</template>

<script>

  import {toTime} from '../../lib/time';

  export default {

    name: 'schedule',

    props: ['schedule'],

    data() {

      return {

        days: ['M', 'T', 'W', 'R', 'F'],

        colors: ['#0D47A1', '#B71C1C', '#1B5E20', '#E65100', '#4A148C', '#263238']

      }

    },

    computed: {

      coursesByDay() {

        let coursesByDay = {M: [], T: [], W: [], R: [], F: []};

        let start = this.schedule.start-30;
        let length = this.schedule.length+60;

        this.schedule.courses.forEach(({name, meets}, i) => {

            meets.forEach(({days, start_time, end_time, location, instructor}) => {

              if (location === 'WEB') return;

              let style = {

                top: ((start_time-start)/length*100)+'%',
                height: ((end_time-start_time)/length*100)+'%',
                background: this.colors[i]

              };

              days.split('').forEach(day => {

                coursesByDay[day].push({name, start: toTime(start_time, true), end: toTime(end_time, true), location, instructor: instructor.replace(' (P)', ''), style});

              });

            });

          });

        return coursesByDay;

      },

      scheduleHeight() {

        return {height: ((this.schedule.length+60)*1.2)+'px'};

      }

    }

  }

</script>

<style scoped>

  .schedule {

    border: 1px solid #ddd;

  }

  .schedule-head {

    display: flex;
    padding: 10px;
    background: #eee;
    text-align: center;
    border-bottom: 1px solid #ddd;

    span {

      flex: 1;

    }

  }

  .schedule-body {

    display: flex;
    font-size: .7rem;

  }

  .day {

    flex: 1;
    border-right: 1px solid #eee;
    padding: 10px;
    position: relative;

    &:last-of-type {

      border: none;

    }

  }

  .course {

    position: absolute;
    left: 5px;
    right: 5px;
    background: #333;
    color: #fff;

  }

  .course-meta {

    margin-top: 5px;
    margin-left: 5px;

    span {

      display: block;
      margin-bottom: 3px;

    }

    span:first-of-type {

      font-weight: 900;

    }

  }

</style>
