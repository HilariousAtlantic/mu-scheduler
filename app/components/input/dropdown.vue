<template>

  <div class="dropdown">

    <button @click="toggleDropdown">{{selectedOption}} <i class="fa fa-angle-down"></i></button>

    <ul v-if="showDropdown" class="option-list">

      <li v-for="option in options" @click="selectOption(option)">{{option}}</li>

    </ul>

  </div>

</template>

<script>

  export default {

    name: 'dropdown',

    props: ['options', 'defaultOption'],

    data() {

      return {

        selectedOption: this.defaultOption || this.options[0],

        showDropdown: false

      }

    },

    methods: {

      selectOption(option) {

        this.selectedOption = option;
        this.showDropdown = false;
        this.$emit('select', option);

      },

      toggleDropdown() {

        this.showDropdown = !this.showDropdown;

      }

    }

  }

</script>

<style scoped>

  .dropdown {

    display: inline-block;
    position: relative;

  }

  button {

    background: #fff;
    border: 1px solid #ddd;
    padding: 10px;
    width: 150px;
    text-align: left;
    outline: none;
    cursor: pointer;

  }

  i {

    float: right;

  }

  ul {

    margin: 0;
    padding: 0;
    width: 148px;
    position: absolute;
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
