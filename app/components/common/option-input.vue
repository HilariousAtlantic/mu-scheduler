<template>

  <select class="option-input form-select" :value="selectedOption" @change="handleOptionChange">

    <option v-for="option in options">{{renderOption(option)}}</option>

  </select>

</template>

<script>

  export default {

    name: 'option-input',

    props: ['options', 'defaultValue', 'renderOption'],

    data() {

      return {

        selectedOption: this.defaultValue || this.options[0]

      }

    },

    methods: {

      handleOptionChange(event) {

        this.selectedOption = event.target.value;

        this.$emit('change', this.selectedOption);

      }

    }

  }

</script>

<style scoped>

</style>
