<template>

  <form class="course-search">

    <select class="form-select" @change="selectTerm">

      <option
        v-for="term in $store.state.terms"
        :value="JSON.stringify(term)"
      >{{term.name}}</option>

    </select>

    <div class="search">

      <input
        type="text"
        class="form-control input-block"
        placeholder="Search Courses"
        v-model="searchTerm"
        @focus="focused = true"
        @blur="focused = false"
      >

      <div v-if="showResults" class="search-results">

        <div
          class="search-result"
          v-for="course in searchResults"
          @mousedown="selectCourse(course)"
        >{{course.searchableName}}</div>

      </div>

    </div>

  </form>

</template>

<script>

  export default {

    name: 'course-search',

    data() {

      return {

        searchTerm: '',
        focused: false

      }

    },

    computed: {

      searchResults() {

        const trim = term => term.trim().toLowerCase().replace(/\W+/g, '');

        let term = trim(this.searchTerm);

        return this.$store.state.courses.filter(course =>
          trim(course.searchableName).indexOf(term) != -1
        ).sort(({searchableName: a}, {searchableName: b}) => {
          let compare = trim(a).indexOf(term)-trim(b).indexOf(term);
          return compare == 0 ? a.localeCompare(b) : compare;
        });

      },

      showResults() {

        return this.focused && this.searchTerm.trim().length > 0;

      }

    },

    methods: {

      selectTerm(event) {

        let term = JSON.parse(event.target.value);
        this.$store.dispatch('selectTerm', term);

      },

      selectCourse(course) {

        this.$store.dispatch('selectCourse', course);
        this.searchTerm = '';

      }

    }

  }

</script>

<style scoped>

  .course-search {

    display: flex;

  }

  .search-results {

    position: absolute;
    top: calc(100% + 5px);
    width: 100%;
    font-size: 12px;
    color: #586069;
    background-color: #fff;
    background-clip: padding-box;
    border: 1px solid rgba(27,31,35,0.15);
    border-radius: 3px;
    box-shadow: 0 3px 12px rgba(27,31,35,0.15);
    z-index: 10;
    max-height: 200px;
    overflow-y: scroll;
  }

  .results-header {

    padding: 8px 10px;
    line-height: 16px;
    background: #f6f8fa;
    border-bottom: 1px solid #e1e4e8;

  }

  .search-result {

    display: block;
    padding: 8px;
    overflow: hidden;
    color: inherit;
    cursor: pointer;
    border-bottom: 1px solid #eaecef;

  }

  .search {

    flex: 1;
    position: relative;

  }

  .form-select {

    margin-right: 5px;
    padding-right: 50px;

  }

</style>
