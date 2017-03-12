<template>

  <div class="courses-view">

    <course-search></course-search>

    <div v-if="showInstructions" class="blankslate">
      <h2>No Courses Selected</h2>
      <p>Use the search bar to select the classes you would like to take</p>
    </div>

    <div v-else>

      <course-list></course-list>

      <button
        class="btn btn-block btn-primary"
        @click="generateSchedules">

        <span>Generate Schedules</span>

      </button>

    </div>

  </div>

</template>

<script>

  import CourseSearch from '../components/course/course-search.vue';
  import CourseList from '../components/course/course-list.vue';

  export default {

    name: 'courses-view',

    components: {CourseSearch, CourseList},

    computed: {

      showInstructions() {

        return this.$store.state.selectedCourses.length === 0;

      }

    },

    methods: {

      generateSchedules() {

        this.$store.dispatch('generateSchedules').then(schedules => {

          if (schedules) this.$router.push('schedules');

        });

      }

    }

  }

</script>

<style scoped>

  .blankslate {
    padding: 50px 30px;
  }

  .course-search, .course-list {
    margin-bottom: 10px;
  }

  .btn-primary {
    padding: 10px;
  }

</style>
