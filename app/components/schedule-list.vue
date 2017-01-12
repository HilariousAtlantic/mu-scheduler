<template>

  <div class="schedule-list">

    <div class="schedule-toolbar">

      <div class="schedule-search">

        <button><i class="fa fa-angle-left"></i></button>
        <span>{{index}} of {{length}}</span>
        <button><i class="fa fa-angle-right"></i></button>

      </div>

      <button class="schedule-select">Save Schedule</button>

    </div>

    <schedule></schedule>

  </div>

</template>

<script>

  import Schedule from './schedule.vue';

  export default {

    name: 'schedule-list',

    components: {Schedule},

    data() {

      return {

        index: 1,

        length: 25

      }

    }

  }

</script>

<style lang="postcss" scoped>

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

  }

  .schedule {

    flex: 1;

  }

</style>
