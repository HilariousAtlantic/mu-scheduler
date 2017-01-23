<template>

  <div class="course-search">

    <button type="button" @click="toggleTermList">{{selectedTermName}}<i class="fa fa-angle-down"></i></button>

    <ul class="term-list" v-if="showTermList">

      <li v-for="term in $store.state.terms" @click="selectTerm(term)">{{term.name}}</li>

    </ul>

    <search-input :options="courses" :limit="5" @select="selectCourse"></search-input>

  </div>

</template>

<script>

  import SearchInput from '../input/search-input.vue';

  export default {

    name: 'course-search',

    components: {SearchInput},

    data() {

      return {

        showTermList: false,

        selectedTerm: {name: 'Select Term'}

      }

    },

    computed: {

      selectedTermName() {

        return this.$store.state.selectedTerm.name || 'Select Term';

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
        this.showTermList = false;

      },

      toggleTermList() {

        this.showTermList = !this.showTermList;

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

  button {

    background: #eee;
    border: 1px solid #ddd;
    width: 250px;
    margin-right: 10px;
    padding: 10px;
    text-align: left;
    cursor: pointer;
    outline: none;

    i {

      float: right;

    }

  }

  .term-list {

    border: 1px solid #ddd;
    border-top: none;
    top: 100%;
    left: 0;
    margin: 0;
    padding: 0;
    position: absolute;
    width: 248px;

  }

  li {

    padding: 10px;
    cursor: pointer;
    background: #fff;
    list-style: none;

    &:hover {

      background: #f5f5f5;

    }

  }

  .search-input {

    flex: 1;

  }

</style>
