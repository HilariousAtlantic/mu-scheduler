<template>

  <input
    type="number"
    class="number-input form-control"
    v-model="number"
    @change="handleNumberChange"
  />

</template>

<script>

  export default {

    name: 'number-input',

    props: ['defaultValue'],

    data() {

      return {

        number: this.defaultValue || 1

      }

    },

    methods: {

      handleNumberChange() {

        this.$emit('change', this.number);

      }

    }

  }

</script>

<style scoped>

  input {
    width: 50px;
  }

</style>
