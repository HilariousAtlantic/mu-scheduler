<template>

  <div class="course-list">

    <div class="course-search">

      <option-input
        :options="$store.state.terms"
        :selectedOption="$store.state.selectedTerm"
        :renderOption="term => term.name || 'Select a Term'"
        @select="selectTerm"
      ></option-input>

      <search-input
        :placeholder="'Search Courses'"
        :options="$store.state.courses"
        :limit="5"
        :renderResult="course => course.searchableName"
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

    methods: {

      selectTerm(term) {

        this.$store.dispatch('selectTerm', term);

      },

      selectCourse(course) {

        this.$store.dispatch('selectCourse', course);

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
