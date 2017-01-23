<template>

  <div class="number-input">

    <span>{{number}}</span>

    <div class="stepper">

      <i class="fa fa-caret-up" @click="incrementNumber"></i>

      <i class="fa fa-caret-down" @click="decrementNumber"></i>

    </div>

  </div>

</template>

<script>

  export default {

    name: 'number-input',

    props: ['step', 'defaultNumber'],

    data() {

      return {

        number: this.defaultNumber || 1

      }

    },

    methods: {

      incrementNumber() {

        this.number += this.step;

        this.$emit('change', this.number);

      },

      decrementNumber() {

        this.number -= this.step;

        this.$emit('change', this.number);

      }

    }

  }

</script>

<style scoped>

  .number-input {

    display: inline-flex;

  }

  span {

    border: 1px solid #ddd;
    padding: 10px 20px;

  }

  .stepper {

    display: flex;
    flex-direction: column;
    justify-content: space-between;
    background: #eee;
    border: 1px solid #ddd;
    border-left: none;

  }

  i {

    flex: 1;
    padding: 0px 5px;
    cursor: pointer;

    &:hover {

      background: #ddd;

    }

  }

</style>
