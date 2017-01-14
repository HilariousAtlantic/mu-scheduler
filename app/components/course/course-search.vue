<template>

  <div class="course-search">

    <button type="button" @click="toggleTermList">{{selectedTerm}}<i class="fa fa-caret-down"></i></button>

    <ul class="term-list" v-if="showTermList">

      <li v-for="term in sortedTerms" @click="selectTerm(term)">{{term.name}}</li>

    </ul>

    <input type="text" placeholder="Filter Courses" autocomplete="off" @input="setFilter" />

  </div>

</template>

<script>

  export default {

    name: 'course-search',

    data() {

      return {

        showTermList: false

      }

    },

    computed: {

      sortedTerms() {

        let sorted = [];
        let seasons = ['Winter', 'Spring', 'Summer', 'Fall'];
        let termsByYear = {};

        this.$store.state.terms.forEach(term => {

          let year = parseInt(term.name.split(' ')[2]);

          if (termsByYear[year]) {
            termsByYear[year].push(term)
          } else {
            termsByYear[year] = [term];
          }

        });

        Object.keys(termsByYear).forEach(year => {
          termsByYear[year] = termsByYear[year].sort((a, b) => {
            return seasons.indexOf(b.name.split(' ')[0]) - seasons.indexOf(a.name.split(' ')[0]);
          });
        });

        Object.keys(termsByYear).sort((a, b) => b - a).forEach(year => {
          sorted = sorted.concat(termsByYear[year]);
        })

        return sorted;

      },

      selectedTerm() {

        let term = this.$store.state.selectedTerm.name;

        return term ? term : 'Select Term';

      }

    },

    methods: {

      selectTerm(term) {

        this.$store.dispatch('selectTerm', term);
        this.showTermList = false;

      },

      setFilter(event) {

        this.$store.dispatch('setFilter', event.target.value);

      },

      toggleTermList() {

        this.showTermList = !this.showTermList;

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

  input {

    flex: 1;
    border: 1px solid #ddd;
    font-size: 1rem;
    padding: 10px;
    outline: none;

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

</style>
