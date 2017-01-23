<template>

  <div class="courses-view">

    <double-header>

      <span slot="left">Selected Courses {{selectedTerm}}</span>

      <span slot="right">{{totalCredits}} Credits</span>

    </double-header>

    <course-search></course-search>

    <selected-courses></selected-courses>

    <router-link to="/filters">

      <button type="button" class="schedules-generate" @click="generateSchedules">Generate Schedules</button>

    </router-link>

  </div>

</template>

<script>

  import DoubleHeader from '../components/common/double-header.vue';
  import CourseSearch from '../components/course/course-search.vue';
  import SelectedCourses from '../components/course/selected-courses.vue';

  export default {

    name: 'courses-view',

    components: {DoubleHeader, CourseSearch, SelectedCourses},

    computed: {

      selectedTerm() {

        let term = this.$store.state.selectedTerm.name;

        return term ? 'for ' + term : '';

      },

      totalCredits() {

        let minTotal = 0;
        let maxTotal = 0;

        this.$store.getters.selectedCourses.forEach(({credits}) => {

          let [min, max] = credits.split('-');

          minTotal += parseInt(min);

          maxTotal += parseInt(max ? max : min);

        });

        return minTotal === maxTotal ? minTotal : minTotal + '-' + maxTotal;

      }

    },

    methods: {

      generateSchedules() {

        this.$store.dispatch('generateSchedules');

      }

    }

  }

</script>

<style scoped>

  .courses-view {

    width: 90%;
    max-width: 1000px;
    margin: 20px auto;

  }

  .double-header {

    font-weight: 900;
    font-size: 1.25rem;

  }

  .course-search {

    margin: 20px 0;

  }

  .selected-courses {

    margin-bottom: 20px;

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
