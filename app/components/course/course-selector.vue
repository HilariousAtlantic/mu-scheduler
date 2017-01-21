<template>

  <div class="course-selector">

    <course-search></course-search>

    <selected-courses></selected-courses>

    <router-link to="/schedules">

      <button class="schedules-generate" type="button" @click="handleGenerateSchedules">Generate Schedules</button>

    </router-link>

  </div>

</template>

<script>

  import CourseSearch from './course-search.vue';
  import SelectedCourses from './selected-courses.vue';

  export default {

    name: 'course-selector',

    components: {CourseSearch, SelectedCourses},

    computed: {

      selectedTerm() {

        let term = this.$store.state.selectedTerm.name;

        return term ? ' for ' + term : '';

      },

      filteredCourses() {

        let trim = (term) => term.toLowerCase().replace(/\W+/g, '');

        return this.$store.state.courses.filter(({subject, number, name}) =>
          trim(subject+number+name).indexOf(trim(this.$store.state.coursesFilter)) != -1
        );

      }

    },

    methods: {

      handleGenerateSchedules() {

        this.$store.dispatch('generateSchedules');

      }

    }

  }

</script>

<style scoped>

  .course-selector {

    display: flex;
    flex-direction: column;
    font-family: sans-serif;
    color: #333;

  }

  .selected-courses {

    flex: 1;
    margin: 10px 0;

  }

  .schedules-generate {

    width: 100%;
    padding: 15px;
    border: 1px solid #ddd;
    background: #4CAF50;
    color: #fff;
    outline: none;
    cursor: pointer;

  }

</style>
