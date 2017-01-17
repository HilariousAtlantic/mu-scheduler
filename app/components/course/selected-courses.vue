<template>

  <div class="selected-courses">

    <header>Selected Courses {{selectedTerm}}</header>

    <ul>

      <li v-for="course in selectedCourses" @click="onCourseClick(course)">

        <detailed-course :course="course"></detailed-course>

      </li>

    </ul>

  </div>



</template>

<script>

  import DetailedCourse from './detailed-course.vue';

  export default {

    name: 'selected-courses',

    components: {DetailedCourse},

    methods: {

      onCourseClick(course) {

        this.$store.dispatch('deselectCourse', course);

      }

    },

    computed: {

      selectedTerm() {

        let term = this.$store.state.selectedTerm.name;

        return term ? 'for ' + term : '';

      },

      selectedCourses() {

        return this.$store.getters.selectedCourses.map(course => {

          let sections = course.sections.map(({name, meets}) => {

            let instructors = [];

            meets.sort((a, b) => {

              if (a.instructor.indexOf('(P)') != -1) {

                return -1;

              } else if (b.instructor.indexOf('(P)') != -1) {

                return 1;

              }

              return 0;

            }).forEach(({instructor}) => {

              if (instructors.indexOf(instructor) == -1) {

                instructors.push(instructor.replace(' (P)', ''));

              }

            });

            return {name, instructor: instructors.join(', ')};

          }).sort((a, b) => a.name.localeCompare(b.name));

          return Object.assign({}, course, {sections});

        })

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

    background: #eee;
    border: 1px solid #ddd;
    padding: 10px;

  }

  ul {

    padding: 0;
    margin: 0;
    flex: 1;
    overflow-y: auto;
    border: 1px solid #ddd;
    border-top: none;

  }

  li {

    list-style: none;
    padding: 10px;
    cursor: pointer;

    &:hover {

      background: #f5f5f5;

      .course-info {

        display: inline-block;

      }

    }

  }

  .course-info {

    display: none;
    float: right;

  }

  .loader-wrapper {

    flex: 1;
    display: flex;
    justify-content: center;
    align-items: center;
    border: 1px solid #ddd;
    border-top: none;

  }

  .loader {

    width: 50px;
    height: 50px;
    border: 5px solid #ddd;
    border-bottom: 5px solid #444;
    border-radius: 90px;
    animation: spin 1s linear infinite;

  }

  @keyframes spin {

    to {

      transform: rotate(360deg);

    }

  }

</style>
