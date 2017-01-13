<template>

  <div class="schedule-list">

    <div class="schedule-toolbar">

      <div class="schedule-search">

        <button @click="decreaseIndex"><i class="fa fa-angle-left"></i></button>
        <span>{{index+1}} of {{schedules.length}}</span>
        <button @click="increaseIndex"><i class="fa fa-angle-right"></i></button>

      </div>

      <button class="schedule-select">Save Schedule</button>

    </div>

    <schedule :courses="currentSchedule.courses"></schedule>

  </div>

</template>

<script>

  import Schedule from './schedule.vue';

  export default {

    name: 'schedule-list',

    components: {Schedule},

    data() {

      return {

        index: 0,

        schedules: [

          {

            courses: [

              {name: 'CSE 102 A', meets: [{days: 'MW', start_time: '13:00', end_time: '13:55', location: 'BEN 102'}, {days: 'F', start_time: '12:00', end_time: '13:50', location: 'BEN 010'}]},

              {name: 'CSE 262 B', meets: [{days: 'MW', start_time: '14:30', end_time: '15:50', location: 'BEN 213'}]},

              {name: 'CSE 278 C', meets: [{days: 'TR', start_time: '10:00', end_time: '11:50', location: 'BEN 010'}]},

              {name: 'CSE 385 D', meets: [{days: 'TR', start_time: '08:30', end_time: '09:50', location: 'BEN 016'}]},

              {name: 'PHY 192 E', meets: [{days: 'MWF', start_time: '10:00', end_time: '11:50', location: 'KRG 311'}]},

              {name: 'ECO 202H F', meets: [{days: 'TR', start_time: '14:30', end_time: '15:50', location: 'BEN 102'}]}

            ]

          },

          {

            courses: [

              {name: 'CSE 102 F', meets: [{days: 'MWF', start_time: '13:00', end_time: '13:55', location: 'BEN 102'}]},

              {name: 'CSE 262 E', meets: [{days: 'MW', start_time: '14:30', end_time: '15:50', location: 'BEN 213'}]},

              {name: 'CSE 278 D', meets: [{days: 'TR', start_time: '10:00', end_time: '11:50', location: 'BEN 010'}]},

              {name: 'CSE 385 C', meets: [{days: 'TR', start_time: '08:30', end_time: '09:50', location: 'BEN 016'}]},

              {name: 'PHY 192 B', meets: [{days: 'MWF', start_time: '10:00', end_time: '11:50', location: 'KRG 311'}]},

              {name: 'ECO 202H A', meets: [{days: 'TR', start_time: '14:30', end_time: '15:50', location: 'BEN 102'}]}

            ]

          }

        ]

      }

    },

    computed: {

      currentSchedule() {

        return this.schedules[this.index];

      }

    },

    methods: {

      increaseIndex() {

        this.index = this.index+1 >= this.schedules.length ?
          this.schedules.length-1 : this.index+1;

      },

      decreaseIndex() {

        this.index = this.index-1 < 0 ?
          0 : this.index-1;

      }

    }

  }

</script>

<style scoped>

  .schedule-list {

    display: flex;
    flex-direction: column;

  }

  .schedule-toolbar {

    margin-bottom: 10px;

  }

  .schedule-search {

    display: inline-block;

    button {

      border: 1px solid #ddd;
      background: #eee;
      padding: 10px 50px;
      outline: none;
      cursor: pointer;

    }

    span {

      text-align: center;
      border: 1px solid #ddd;
      padding: 10px 75px;

    }

  }

  .schedule-select {

    float: right;
    display: inline-block;
    padding: 10px 100px;
    border: 1px solid #ddd;
    background: #66BB6A;
    color: #fff;
    outline: none;
    cursor: pointer;

  }

  .schedule {

    flex: 1;

  }

</style>
