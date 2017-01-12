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
              <span>{{course.start_time}} - {{course.end_time}}</span><br>
              <span>{{course.location}}</span>

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

        colors: ['#283593', '#B71C1C', '#00695C', '#FF6F00', '#4A148C', '#263238']

      }

    },

    computed: {

      coursesByDay() {

        let start = 420; //7am
        let end = 1320; //10pm

        function getStartPercentage(startTime) {

          let [hours, minutes] = startTime.split(':');
          let into = parseInt(hours)*60+parseInt(minutes);

          return (into-start)/(end-start)*100;

        }

        function getLengthPercentage(startTime, endTime) {

          let [startHours, startMinutes] = startTime.split(':');
          let [endHours, endMinutes] = endTime.split(':');
          let length = parseInt(endHours)*60+parseInt(endMinutes)-parseInt(startHours)*60-parseInt(startMinutes);

          return length/(end-start)*100;

        }

        let coursesByDay = {M: [], T: [], W: [], R: [], F: [], S: []};

        this.courses
          .map((course, i) => Object.assign(course, {color: this.colors[i] || "#000"}))
          .forEach(({name, meets, color}) => {

            meets.forEach(({days, start_time, end_time, location}) => {

              days.split('').forEach(day => {

                let style = {

                  top: getStartPercentage(start_time)+'%',
                  height: getLengthPercentage(start_time, end_time)+'%',
                  background: color,
                  color: '#fff'

                }

                coursesByDay[day].push({name, start_time, end_time, location, style});

              })

            })

          })

        return coursesByDay;

      }

    }

  }

</script>

<style lang="postcss" scoped>

  .schedule {

    border: 1px solid #ddd;
    font-size: .75rem;
    display: flex;
    flex-direction: column;

  }

  .schedule-head {

    display: flex;
    padding: 10px;
    background: #eee;
    text-align: center;

    span {

      flex: 1;

    }

  }

  .schedule-body {

    flex: 1;
    display: flex;
    overflow-y: auto;

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
    background: #eee;

  }

  .course-meta {

    margin-top: 5px;
    margin-left: 5px;

  }

</style>
