<template>

  <div class="selected-courses">

    <header>

      <span class="title">Selected Courses {{selectedTerm}}</span>

      <span class="credits">{{totalCredits}} Credits</span>

    </header>

    <detailed-course v-for="course in $store.getters.selectedCourses" :course="course"></detailed-course>

  </div>

</template>

<script>

  import DetailedCourse from './detailed-course.vue';

  export default {

    name: 'selected-courses',

    components: {DetailedCourse},

    methods: {

      deselectCourse(course) {

        this.$store.dispatch('deselectCourse', course);

      }

    },

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

    }

  }

</script>

<style scoped>

  .selected-courses {

    display: flex;
    flex-direction: column;

  }

  header {

    display: flex;
    margin-bottom: 20px;
    font-size: 1.25rem;
    font-weight: 900;

  }

  .title {

    flex: 1;

  }

</style>
