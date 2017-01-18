<template>

  <div class="schedule">

    <div class="schedule-head">

      <span v-for="day in days">{{day}}</span>

    </div>

    <div class="schedule-body">

      <div class="day" v-for="day in days">

          <div class="course" v-for="course in coursesByDay[day]" :style="course.style">

            <div class="course-meta">

              <span>{{course.name}}</span><br>
              <span>{{course.start}} - {{course.end}}</span><br>
              <span>{{course.location}}</span><br>
              <span>{{course.instructor}}</span>

            </div>

          </div>

      </div>

    </div>

  </div>

</template>

<script>

  export default {

    name: 'schedule',

    props: ['courses'],

    data() {

      return {

        days: ['M', 'T', 'W', 'R', 'F', 'S'],

        colors: ['#0D47A1', '#B71C1C', '#1B5E20', '#E65100', '#4A148C', '#263238']

      }

    },

    computed: {

      coursesByDay() {

        let coursesByDay = {M: [], T: [], W: [], R: [], F: [], S: []};

        let start = 420; //7am
        let end = 1320; //10pm

        function formatTime(time) {

          let hours = Math.floor(time/60);
          let minutes = time-hours*60;

          if (hours > 12) {
            hours -= 12;
          }

          return hours + ':' + ('00'+minutes).slice(-2);

        }

        this.courses.forEach(({name, meets}, i) => {

            meets.forEach(({days, start_time, end_time, location, instructor}) => {

              if (location === 'WEB') return;

              let style = {

                top: ((start_time-start)/(end-start)*100)+'%',
                height: ((end_time-start_time)/(end-start)*100)+'%',
                background: this.colors[i]

              };

              days.split('').forEach(day => {

                coursesByDay[day].push({name, start: formatTime(start_time), end: formatTime(end_time), location, instructor: instructor.replace(' (P)', ''), style});

              });

            });

          });

        return coursesByDay;

      }

    }

  }

</script>

<style scoped>

  .schedule {

    border: 1px solid #ddd;
    display: flex;
    flex-direction: column;

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

    flex: 1;
    display: flex;
    overflow-y: auto;
    font-size: .75rem;

  }

  .day {

    flex: 1;
    height: 60rem;
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

  }

</style>
