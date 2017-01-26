<template>

  <div class="courses-view">

    <instruction>

      <div>

        <h3>Selecting Courses</h3>

        <p>Welcome to the unofficial course schedule for Miami University! The first thing you need to do is select the courses you would like to take. Choose the term from the dropdown menu, then use the search bar to find the course.</p>

        <p>Once you select all of the courses you would like to take, click on the generate schedules button. This will take you to a page where you can create filters to narrow down the number of schedules. If you would like to return to this page, click on the courses tab at the top of the page.</p>

      </div>

    </instruction>

    <double-header>

      <span slot="left">Selected Courses {{selectedTerm}}</span>

      <span slot="right">{{totalCredits}} Credits</span>

    </double-header>

    <course-list></course-list>

    <router-link to="filters">

      <button type="button" class="schedules-generate" @click="generateSchedules">Generate Schedules</button>

    </router-link>

  </div>

</template>

<script>

  import Instruction from '../components/common/instruction.vue';
  import DoubleHeader from '../components/common/double-header.vue';
  import CourseList from '../components/course/course-list.vue';

  export default {

    name: 'courses-view',

    components: {Instruction, DoubleHeader, CourseList},

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

  .instruction {

    margin-bottom: 10px;

    h3 {

      margin-top: 0;

    }

    p {

      margin-top: 0;
      margin-bottom: 10px;

      &:last-of-type {

        margin-bottom: 0;

      }

    }

  }

  .double-header {

    font-weight: 900;
    font-size: 1.25rem;

  }

  .course-list {

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

    &:hover {

      background: #43A047;

    }

  }

</style>
