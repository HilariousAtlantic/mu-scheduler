<template>

  <div class="courses-view">

    <div class="flash mb-3">

      <p>

        This is an early version of the application, so there will be some
        problems and not all of the features are added or completed. If you run
        into any issues, feel free to email us at

        <strong>hilariousatlantic@gmail.com</strong>

      </p>

    </div>

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

        this.$store.dispatch('generateSchedules').then(success => {

          if (success) this.$router.push('schedules');

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
