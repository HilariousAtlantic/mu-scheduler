<template>

  <div class="courses-view">

    <course-search></course-search>

    <course-list></course-list>

    <button
      class="btn btn-block btn-primary"
      :class="{disabled: !canGenerateSchedules}"
      @click="generateSchedules">

      <span>Generate Schedules</span>

    </button>

  </div>

</template>

<script>

  import CourseSearch from '../components/course/course-search.vue';
  import CourseList from '../components/course/course-list.vue';

  export default {

    name: 'courses-view',

    components: {CourseSearch, CourseList},

    computed: {

      canGenerateSchedules() {

        return this.$store.state.selectedCourses.length > 0;

      }

    },

    methods: {

      generateSchedules() {

        if (this.canGenerateSchedules) {

          this.$store.dispatch('generateSchedules').then(schedules => {

            if (schedules) this.$router.push('schedules');

          });

        }

      }

    }

  }

</script>

<style scoped>

  .course-list {
    margin: 10px 0;
  }

  .btn-primary {
    padding: 10px;
  }

</style>
