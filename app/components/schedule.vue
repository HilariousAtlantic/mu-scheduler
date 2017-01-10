<template>

  <div class="schedule">

    <div class="day" v-for="day in days">

      <div class="day-head">{{day}}</div>

      <div class="day-body">

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

    data() {

      return {

        days: ['M', 'T', 'W', 'R', 'F', 'S'],

        courses: [

          {name: 'CSE 102', meets: [{days: 'MW', start_time: '13:00', end_time: '13:55', location: 'BEN 102'}, {days: 'F', start_time: '12:00', end_time: '13:50', location: 'BEN 010'}]},

          {name: 'CSE 262', meets: [{days: 'MW', start_time: '14:30', end_time: '15:50', location: 'BEN 213'}]},

          {name: 'CSE 278', meets: [{days: 'TR', start_time: '10:00', end_time: '11:50', location: 'BEN 010'}]},

          {name: 'CSE 385', meets: [{days: 'TR', start_time: '08:30', end_time: '09:50', location: 'BEN 016'}]},

          {name: 'PHY 192', meets: [{days: 'MWF', start_time: '10:00', end_time: '11:50', location: 'KRG 311'}]},

          {name: 'ECO 202H', meets: [{days: 'TR', start_time: '14:30', end_time: '15:50', location: 'BEN 102'}]}

        ]

      }

    },

    computed: {

      coursesByDay() {

        let start = 450; //7:30am
        let end = 1080; //6pm

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

        this.courses.forEach(({name, meets}) => {

          meets.forEach(({days, start_time, end_time, location}) => {

            days.split('').forEach(day => {

              let style = {

                top: getStartPercentage(start_time)+'%',
                height: getLengthPercentage(start_time, end_time)+'%'

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

    display: flex;
    border: 1px solid #ddd;
    font-size: .75rem;

  }

  .day {

    flex: 1;
    display: flex;
    flex-direction: column;
    border-right: 1px solid #eee;

    &:last-of-type {

      border: none;

    }

  }

  .day-head {

    padding: 10px;
    text-align: center;
    background: #eee;
    border-bottom: 1px solid #ddd;

  }

  .day-body {

    flex: 1;
    padding: 10px;
    position: relative;
    overflow-y: auto;

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
