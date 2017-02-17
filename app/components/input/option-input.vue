<template>

  <div class="option-input">

    <button @click="showDropdown = true" @blur="showDropdown = false">

      {{renderOption(selectedOption)}} <i class="fa fa-angle-down"></i>

    </button>

    <ul v-if="showDropdown" class="option-list">

      <li v-for="option in options" @mousedown="selectOption(option)">

        <span>{{renderOption(option)}}</span>

      </li>

    </ul>

  </div>

</template>

<script>

  export default {

    name: 'option-input',

    props: ['options', 'selectedOption', 'renderOption'],

    data() {

      return {

        showDropdown: false

      }

    },

    methods: {

      selectOption(option) {

        this.selectedOption = option;
        this.showDropdown = false;
        this.$emit('select', option);

      }

    }

  }

</script>

<style scoped>

  .option-input {

    display: inline-block;
    position: relative;

  }

  button {

    background: #fff;
    border: 1px solid #ddd;
    padding: 10px;
    width: 100%;
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
    width: 100%;
    box-sizing: border-box;
    position: absolute;
    border: 1px solid #ddd;
    border-top: none;
    background: #fff;
    cursor: pointer;
    z-index: 10;

  }

  li {

    list-style: none;
    padding: 10px;

    &:hover {

      background: #f5f5f5;

    }

  }

</style>
