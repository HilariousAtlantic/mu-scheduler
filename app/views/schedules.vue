<template>

  <div class="schedules-view">

    <filter-list></filter-list>

    <schedule-toolbar :index="index" @dec="decreaseIndex" @inc="increaseIndex"></schedule-toolbar>

    <schedule-list :index="index"></schedule-list>

  </div>

</template>

<script>

  import FilterList from '../components/filter/filter-list.vue';
  import ScheduleToolbar from '../components/schedule/schedule-toolbar.vue';
  import ScheduleList from '../components/schedule/schedule-list.vue';

  export default {

    name: 'schedules-view',

    components: {FilterList, ScheduleToolbar, ScheduleList},

    data() {

      return {

        index: 0

      }

    },

    methods: {

      decreaseIndex() {

        if (this.index > 0) {

          this.index--;

        }

      },

      increaseIndex() {

        if (this.index < this.$store.state.schedules.length-1) {

          this.index++;

        }

      }

    }

  }

</script>

<style scoped>

  .schedule-toolbar {
    margin: 10px 0;
  }

</style>
