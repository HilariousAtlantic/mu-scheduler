<template>

  <div class="search-input">

    <input type="text" v-model="term" @focus="focused = true" @blur="focused = false">

    <ul v-if="showSearchResults" class="search-results">

      <li v-for="result in searchResults" @mousedown="selectResult(result)">{{result.text}}</li>

    </ul>

  </div>

</template>

<script>

  export default {

    name: 'search-input',

    props: ['options', 'limit'],

    data() {

      return {

        term: '',

        focused: false

      }

    },

    computed: {

      searchResults() {

        let trim = (text) => text.toLowerCase().replace(/\W+/g, '');

        let results = this.options.filter(({text}) => trim(text).indexOf(trim(this.term)) != -1);

        return results.slice(0, this.limit || 10);

      },

      showSearchResults() {

        return this.term.trim().length > 0 && this.focused;

      }

    },

    methods: {

      selectResult(result) {

        this.$emit('select', result);

        this.term = '';

      }

    }

  }

</script>

<style scoped>

  .search-input {

    display: inline-block;
    position: relative;

  }

  input {

    background: #fff;
    border: 1px solid #ddd;
    padding: 10px;
    outline: none;
    width: 100%;
    box-sizing: border-box;

  }

  ul {

    position: absolute;
    margin: 0;
    padding: 0;
    left: 0;
    right: 0;
    border: 1px solid #ddd;
    border-top: none;
    background: #fff;
    cursor: pointer;

  }

  li {

    list-style: none;
    padding: 10px;

    &:hover {

      background: #f5f5f5;

    }

  }

</style>
