<template>

  <div class="filtered-courses">

    <header>Course List {{selectedTerm}}</header>

    <ul v-if="!$store.state.requestingCourses">

      <li v-for="course in filteredCourses" @click="onCourseClick(course)">

        {{course.subject}} {{course.number}} - {{course.title}}
        <i class="course-info fa fa-info-circle"></i>

      </li>

    </ul>

    <div class="loader-wrapper" v-if="$store.state.requestingCourses">
      <div class="loader"></div>
    </div>

  </div>



</template>

<script>

  export default {

    name: 'filtered-courses',

    methods: {

      onCourseClick(course) {

        this.$store.dispatch('selectCourse', course);

      }

    },

    computed: {

      selectedTerm() {

        let term = this.$store.state.selectedTerm.name;

        return term ? 'for ' + term : '';

      },

      filteredCourses() {

        let trim = (term) => term.toLowerCase().replace(/\W+/g, '');

        return this.$store.state.courses.filter(({subject, number, title}) =>
          trim(subject+number+title).indexOf(trim(this.$store.state.coursesFilter)) != -1
        );

      }

    }

  }

</script>

<style scoped>

  .filtered-courses {

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
