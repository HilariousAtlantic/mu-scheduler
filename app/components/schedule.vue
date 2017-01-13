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

        colors: ['#0D47A1', '#B71C1C', '#1B5E20', '#E65100', '#4A148C', '#263238']

      }

    },

    computed: {

      coursesByDay() {

        let coursesByDay = {M: [], T: [], W: [], R: [], F: [], S: []};

        let startTime = 420; //7am
        let endTime = 1320; //10pm

        this.courses.forEach(({name, meets}, i) => {

            meets.forEach(({days, start_time, end_time, location}) => {

              let [startHours, startMinutes] = start_time.split(':');
              let [endHours, endMinutes] = end_time.split(':');

              let start = startHours*60+startMinutes;
              let length = endHours*60+endMinutes-startHours*60-startMinutes;

              let style = {

                top: ((start-startTime)/(endTime-startTime)*100)+'%',
                height: (length/(endTime-startTime)*100)+'%',
                background: this.colors[i]

              };

              startHours = startHours > 12 ? startHours-12 : startHours;
              endHours = endHours > 12 ? endHours-12 : endHours;

              start_time = startHours + ':' + startMinutes;
              end_time = endHours + ':' + endMinutes;

              days.split('').forEach(day => {

                coursesByDay[day].push({name, start_time, end_time, location, style});

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
    font-size: .75rem;
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
