<template>

  <div class="schedule-calendar">

    <div class="schedule-head">

      <span v-for="day in formattedDays">{{day}}</span>

    </div>

    <div class="schedule-body" :style="{height: length+'px'}">

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

  import {formatTime} from '../../lib/time';
  import {formatDay} from '../../lib/days';

  export default {

    name: 'schedule-calendar',

    props: ['courses', 'start', 'length'],

    data() {

      return {

        days: ['M', 'T', 'W', 'R', 'F'],

        colors: ['#0D47A1', '#B71C1C', '#1B5E20', '#E65100', '#4A148C', '#263238']

      }

    },

    computed: {

      formattedDays() {

        return this.days.map(formatDay);

      },

      coursesByDay() {

        let coursesByDay = {M: [], T: [], W: [], R: [], F: []};

        this.courses.forEach(({name, meets}, i) => {

            meets.forEach(({days, start_time, end_time, location, instructor}) => {

              if (location === 'WEB') return;

              let style = {

                top: ((start_time-this.start)/this.length*100)+'%',
                height: ((end_time-start_time)/this.length*100)+'%',
                background: this.colors[i]

              };

              days.split('').forEach(day => {

                coursesByDay[day].push({location, instructor, style, name,
                  start: formatTime(start_time, 'h:mm'),
                  end: formatTime(end_time, 'h:mm')
                });

              });

            });

          });

        return coursesByDay;

      }

    }

  }

</script>

<style scoped>

  .schedule-calendar {

    border: 1px solid #ddd;

  }

  .schedule-head {

    display: flex;
    padding: 10px;
    font-size: .8rem;
    background: #eee;
    text-align: center;
    border-bottom: 1px solid #ddd;

    span {

      flex: 1;

    }

  }

  .schedule-body {

    display: flex;
    padding: 5px 0;
    font-size: .7rem;

  }

  .day {

    flex: 1;
    position: relative;
    border-right: 1px solid #eee;

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
