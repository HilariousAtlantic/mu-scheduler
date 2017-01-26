<template>

  <div class="course-list">

    <div class="course-search">

      <option-input
        :options="terms"
        :defaultOption="{text: selectedTermName}"
        @select="selectTerm"
      ></option-input>

      <search-input
        :placeholder="'Search Courses'"
        :options="courses"
        :limit="5"
        @select="selectCourse"
      ></search-input>

    </div>

    <course v-for="course in $store.getters.selectedCourses" :course="course"></course>

  </div>

</template>

<script>

  import OptionInput from '../input/option-input.vue';
  import SearchInput from '../input/search-input.vue';
  import Course from './course.vue';

  export default {

    name: 'course-list',

    components: {OptionInput, SearchInput, Course},

    computed: {

      selectedTermName() {

        return this.$store.state.selectedTerm.name || 'Select Term';

      },

      terms() {

        return this.$store.state.terms.map(term => Object.assign({}, term, {text: term.name}));

      },

      courses() {

        let getText = ({subject, number, title}) => subject + ' ' + number + ' - ' + title;

        return this.$store.state.courses.map(

          course => Object.assign({}, course, {text: getText(course)})

        );

      }

    },

    methods: {

      selectTerm(term) {

        this.$store.dispatch('selectTerm', term);

      },

      selectCourse(course) {

        this.$store.dispatch('selectCourse', course);

      },

      deselectCourse(course) {

        this.$store.dispatch('deselectCourse', course);

      }

    }

  }

</script>

<style scoped>

  .course-search {

    display: flex;
    margin-bottom: 10px;

  }

  .option-input {

    width: 250px;
    margin-right: 10px;

  }

  .search-input {

    flex: 1;

  }

</style>
