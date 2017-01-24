<template>

  <div class="course-search">

    <option-input
      :options="terms"
      :defaultOption="{text: selectedTermName}"
      @select="selectTerm"
    ></option-input>

    <search-input :options="courses" :limit="5" @select="selectCourse"></search-input>

  </div>

</template>

<script>

  import OptionInput from '../input/option-input.vue';
  import SearchInput from '../input/search-input.vue';

  export default {

    name: 'course-search',

    components: {OptionInput, SearchInput},

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

      }

    }

  }

</script>

<style scoped>

  .course-search {

    display: flex;
    position: relative;

  }

  .option-input {

    width: 250px;
    margin-right: 10px;

  }

  .search-input {

    flex: 1;

  }

</style>
